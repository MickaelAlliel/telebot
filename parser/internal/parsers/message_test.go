package parsers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"mickaelalliel.com/telebot/parser/internal/test_utils"
)

type expenseTestCase struct {
	Amount   float64
	Category string
	Method   string
}

func assertExpense(t *testing.T, exp *Expense, err error, testCase expenseTestCase) {
	assert.Nil(t, err)
	assert.Equal(t, testCase.Amount, exp.Amount)
	assert.Equal(t, testCase.Category, exp.Category)
	assert.Equal(t, testCase.Method, exp.Method)
	assert.Equal(t, "test_username", exp.OwnerName)
}

var validTestCases = []expenseTestCase{
	{
		Amount:   32.5,
		Category: "food",
		Method:   "cash",
	},
	{
		Amount:   100,
		Category: "כלבו",
		Method:   "אשראי",
	},
}

func TestParseExpenseValid(t *testing.T) {
	for _, testCase := range validTestCases {
		msg := test_utils.NewTestTelegramMessage(fmt.Sprintf("%g\n%s\n%s", testCase.Amount, testCase.Category, testCase.Method))
		exp, err := ParseExpenseMessage(msg)
		assertExpense(t, exp, err, testCase)
	}
}

func TestParseExpenseEmptyValues(t *testing.T) {
	testCase := expenseTestCase{
		Amount:   0.0,
		Category: "",
		Method:   "",
	}
	msg := test_utils.NewTestTelegramMessage(fmt.Sprintf("%g\n%s\n%s", testCase.Amount, testCase.Category, testCase.Method))
	exp, err := ParseExpenseMessage(msg)
	assert.Nil(t, exp)
	assert.NotNil(t, err)
}

func TestParseExpenseEmptyMessage(t *testing.T) {
	msg := test_utils.NewTestTelegramMessage(" ")
	exp, err := ParseExpenseMessage(msg)
	assert.Nil(t, exp)
	assert.NotNil(t, err)
}

func TestParseExpenseIrrelevantMessage(t *testing.T) {
	msg := test_utils.NewTestTelegramMessage("hey\nwhats up?")
	exp, err := ParseExpenseMessage(msg)
	assert.Nil(t, exp)
	assert.NotNil(t, err)
}
