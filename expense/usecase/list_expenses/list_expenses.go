package list_expenses

import (
	"github.com/RenanLourenco/financial-organizer/expense/adapter/entity"
	"gorm.io/gorm"
)

type ListExpense struct{
	Repository *gorm.DB
}


// Execute List expenses by month
// @Summary List expenses by month
// @Description List expenses by month
// @Produce json
// @Tags Expense
// @Success 200 {object} ListExpenseOutput
// @Router /expenses/:year/:month [get]
func (c *ListExpense) Execute() (ListExpenseOutput, error) {
	var expenses []entity.Expense

	c.Repository.Find(&expenses)

	output := ListExpenseOutput{
		Data: expenses,
	}

	return output, nil

}

