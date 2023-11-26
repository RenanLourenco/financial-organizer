package search_expense

import (
	"github.com/RenanLourenco/financial-organizer/expense/adapter/entity"
	"gorm.io/gorm"
)

type SearchExpense struct {
	Repository *gorm.DB
}

func (c *SearchExpense) Execute(input SearchExpenseInput) SearchExpenseOutput {
	var expenses []entity.Expense

	searchString := "%" + input.Search + "%"

	c.Repository.Where("description LIKE ?", searchString).Find(&expenses)

	output := SearchExpenseOutput{
		Data: expenses,
	}

	return output
}
