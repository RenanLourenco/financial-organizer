package update_expense

import (
	"errors"

	expenses "github.com/RenanLourenco/financial-organizer/expenses/adapter/entity"
	"gorm.io/gorm"
)


type UpdateExpense struct {
	Repository *gorm.DB
}

func (c *UpdateExpense) Execute(input UpdateExpenseInput) (UpdateExpenseOutput, error) {
	var expense expenses.Expenses
	c.Repository.First(&expense, input.ID)

	if expense.ID == 0 {
		return UpdateExpenseOutput{}, errors.New("Expense not found")
	}

	c.Repository.Model(&expense).UpdateColumns(input.UpdatedExpense)

	output := UpdateExpenseOutput{
		Model: expense.Model,
		Description: expense.Description,
		Value: expense.Value,
		Date: expense.Date,
	}
	
	return output, nil


}