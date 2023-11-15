package delete_expense

import (
	"github.com/RenanLourenco/financial-organizer/expense/adapter/entity"
	"gorm.io/gorm"
)


type DeleteExpense struct {
	Repository *gorm.DB
}


func (c *DeleteExpense) Execute(input DeleteExpenseInput) (DeleteExpenseOutput, error) {
	var expense entity.Expense

	c.Repository.Delete(&expense, input.ID)

	output := DeleteExpenseOutput{
		Msg: "Expense deleted.",
	}
	
	return output, nil
}