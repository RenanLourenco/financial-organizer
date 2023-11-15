package list_expenses

import (
	"github.com/RenanLourenco/financial-organizer/expense/adapter/entity"
	"gorm.io/gorm"
)

type ListExpense struct{
	Repository *gorm.DB
}



func (c *ListExpense) Execute() (ListExpenseOutput, error) {
	var expenses []entity.Expense

	c.Repository.Find(&expenses)

	output := ListExpenseOutput{
		Data: expenses,
	}

	return output, nil

}

