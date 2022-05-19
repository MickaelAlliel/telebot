package parsers

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"mickaelalliel.com/telebot/parser/internal/utils"
)

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

func isValidMethod(text string) bool {
	switch strings.ToLower(text) {
	case "cash", "check", "card", "credit card", "credit", "wire", "wire transfer", "bank transfer", "bit", "אשראי", "כרטיס", "כרטיס אשראי", "מזומן", "צק", "העברה", "העברה בנקאית", "ביט":
		return true
	default:
		return false
	}
}

func ParseExpenseMessage(message *tgbotapi.Message) (*Expense, error) {
	var amount float64 = 0.0
	var category = ""
	var method = ""

	lines := strings.Split(message.Text, "\n")

	for _, line := range lines {
		if val, err := strconv.ParseFloat(line, 64); err == nil {
			amount = val
			continue
		}
		if isValidMethod(line) {
			method = line
			continue
		}
	}

	category = strings.Replace(message.Text, strconv.FormatFloat(amount, 'f', -1, 64), "", -1)
	category = strings.Replace(category, method, "", -1)
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
