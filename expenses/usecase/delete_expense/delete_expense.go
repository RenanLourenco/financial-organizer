package delete_expense

import (
	expenses "github.com/RenanLourenco/financial-organizer/expenses/adapter/entity"
	"gorm.io/gorm"
)


type DeleteExpense struct {
	Repository *gorm.DB
}


func (c *DeleteExpense) Execute(input DeleteExpenseInput) (DeleteExpenseOutput, error) {
	var expense expenses.Expenses

	c.Repository.Delete(&expense, input.ID)

	output := DeleteExpenseOutput{
		Msg: "Expense deleted.",
	}
	
	return output, nil
}