package update_expense

import (
	"errors"

	"github.com/RenanLourenco/financial-organizer/expense/adapter/entity"
	"gorm.io/gorm"
)


type UpdateExpense struct {
	Repository *gorm.DB
}


// Execute Update specific expense month
// @Summary Update specific expense month
// @Description Update specific expense month
// @Produce json
// @Tags Expense
// @Success 200 {object} func (c *UpdateExpense) Execute(input UpdateExpenseInput) (UpdateExpenseOutput, error) {

// @Router /expenses/:id [put]
func (c *UpdateExpense) Execute(input UpdateExpenseInput) (UpdateExpenseOutput, error) {
	var expense entity.Expense
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