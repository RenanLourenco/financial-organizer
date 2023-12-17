package get_income

import (
	"errors"

	"github.com/RenanLourenco/financial-organizer/income/adapter/entity"
	"gorm.io/gorm"
)

type GetIncome struct {
	Repository *gorm.DB
}

// Execute Get income
// @Summary Get specific income by id
// @Description Get specific income by id
// @Tags Income
// @Success 200 {object} GetIncomeOutput
// @Router /income/:id [get]
func (c *GetIncome) Execute(input GetIncomeInput) (GetIncomeOutput, error){
	var income entity.Income
	c.Repository.First(&income,input.ID)

	if income.ID == 0 {
		return GetIncomeOutput{}, errors.New("Income not found")
	}

	output := GetIncomeOutput{
		Model: income.Model,
		Description: income.Description,
		Value: income.Value,
		Date: income.Date,
	}

	return output, nil


}
