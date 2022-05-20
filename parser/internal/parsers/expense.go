package parsers

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"golang.org/x/exp/slices"
	"mickaelalliel.com/telebot/parser/ent"
	"mickaelalliel.com/telebot/parser/internal/utils"
)

var allowedPaymentMethods = map[string][]string{
	"cash":  {"cash", "מזומן"},
	"card":  {"card", "credit card", "אשראי", "כרטיס", "כרטיס אשראי"},
	"bit":   {"bit", "ביט"},
	"check": {"check", "צק", "צ'ק"},
	"wire":  {"wire", "wire transfer", "bank transfer", "העברה", "העברה בנקאית"},
	"10bis": {"10bis", "תןביס", "תן ביס", "10ביס", "10 ביס"},
	"cibus": {"cibus", "סיבוס"},
}

type Expense struct {
	Amount    float64    `json:"amount"`
	Category  string     `json:"category"`
	Method    string     `json:"method"`
	OwnerName string     `json:"ownerName"`
	Timestamp *time.Time `json:"timestamp"`
}

func (exp *Expense) isValid() bool {
	if exp.Amount <= 0 || exp.Category == "" || exp.Method == "" || exp.OwnerName == "" || exp.Timestamp == nil {
		return false
	}
	return true
}

func (exp *Expense) SaveEnt(ctx context.Context, db *ent.Client) (*ent.Expense, error) {
	e, err := db.Expense.Create().
		SetAmount(exp.Amount).
		SetCategory(exp.Category).
		SetMethod(exp.Method).
		SetOwnerName(exp.OwnerName).
		SetTimestamp(*exp.Timestamp).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	return e, nil
}

func getNormalizedPaymentMethod(text string) string {
	paymentMethod := strings.ToLower(text)
	for key, options := range allowedPaymentMethods {
		if slices.Contains(options, paymentMethod) {
			return key
		}
	}
	return ""
}

func ParseExpenseMessage(message *tgbotapi.Message) (*Expense, error) {
	var amount float64 = 0.0
	var category = ""
	var method = ""

	lines := strings.Split(message.Text, "\n")
	originalPaymentMethod := ""

	for _, line := range lines {
		if val, err := strconv.ParseFloat(line, 64); err == nil {
			amount = val
			continue
		}
		if paymentMethod := getNormalizedPaymentMethod(line); paymentMethod != "" {
			method = paymentMethod
			originalPaymentMethod = line
			continue
		}
	}

	category = strings.Replace(message.Text, strconv.FormatFloat(amount, 'f', -1, 64), "", -1)
	category = strings.Replace(category, originalPaymentMethod, "", -1)
	category = strings.TrimSpace(category)

	e := &Expense{
		Amount:    amount,
		Category:  category,
		Method:    method,
		OwnerName: message.From.UserName,
		Timestamp: utils.ToTimePtr(message.Time()),
	}

	if !e.isValid() {
		return nil, fmt.Errorf("expense message is invalid, missing fields: %v", e)
	}

	return e, nil
}
