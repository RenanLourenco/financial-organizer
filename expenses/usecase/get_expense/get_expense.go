package get_expense

import (
	"errors"

	expenses "github.com/RenanLourenco/financial-organizer/expenses/adapter/entity"
	"gorm.io/gorm"
)

type GetExpense struct {
	Repository *gorm.DB
}

func (c *GetExpense) Execute (input GetExpenseInput) (GetExpenseOutput, error) {
	var expense expenses.Expenses
	c.Repository.First(&expense, input.ID)

	if expense.ID == 0 {
		return GetExpenseOutput{}, errors.New("Expense not found")
	}

	output := GetExpenseOutput{
		Model: expense.Model,
		Description: expense.Description,
		Value: expense.Value,
		Date: expense.Date,
	}

	return output, nil


}
