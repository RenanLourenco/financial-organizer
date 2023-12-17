package list_incomes

import (
	"github.com/RenanLourenco/financial-organizer/income/adapter/entity"
	"gorm.io/gorm"
)

type ListIncomes struct {
	Repository *gorm.DB
}

// Execute Get incomes
// @Summary List incomes
// @Description List incomes
// @Tags Income
// @Success 200 {object} ListIncomesOutput
// @Router /incomes [get]
func (c *ListIncomes) Execute() (ListIncomesOutput, error){
	var incomes []entity.Income

	c.Repository.Find(&incomes)

	output := ListIncomesOutput{
		Data: incomes,
	}

	return output, nil
}