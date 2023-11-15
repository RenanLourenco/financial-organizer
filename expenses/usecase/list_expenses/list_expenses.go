package list_expenses

import (
	expenses "github.com/RenanLourenco/financial-organizer/expenses/adapter/entity"
	"gorm.io/gorm"
)

type ListExpense struct{
	Repository *gorm.DB
}



func (c *ListExpense) Execute() (ListExpenseOutput, error) {
	var expenses []expenses.Expenses

	c.Repository.Find(&expenses)

	output := ListExpenseOutput{
		Data: expenses,
	}

	return output, nil

}

