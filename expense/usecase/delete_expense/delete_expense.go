package delete_expense

import (
	"github.com/RenanLourenco/financial-organizer/expense/adapter/entity"
	"gorm.io/gorm"
)


type DeleteExpense struct {
	Repository *gorm.DB
}

// Execute Delete specific expense
// @Summary Delete specific expense
// @Description Delete specific expense
// @Produce json
// @Tags Expense
// @Success 200 {object} DeleteExpenseOutput
// @Router /expense/:id [delete]
func (c *DeleteExpense) Execute(input DeleteExpenseInput) (DeleteExpenseOutput, error) {
	var expense entity.Expense

	c.Repository.Delete(&expense, input.ID)

	output := DeleteExpenseOutput{
		Msg: "Expense deleted.",
	}
	
	return output, nil
}