package get_expense

import (
	"errors"

	"github.com/RenanLourenco/financial-organizer/expense/adapter/entity"
	"gorm.io/gorm"
)

type GetExpense struct {
	Repository *gorm.DB
}

func (c *GetExpense) Execute (input GetExpenseInput) (GetExpenseOutput, error) {
	var expense entity.Expense
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
