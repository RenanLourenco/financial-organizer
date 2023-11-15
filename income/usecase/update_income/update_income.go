package update_income

import (
	"errors"

	"github.com/RenanLourenco/financial-organizer/income/adapter/entity"
	"gorm.io/gorm"
)

type UpdateIncome struct {
	Repository *gorm.DB
}

func (c *UpdateIncome) Execute(input UpdateIncomeInput) (UpdateIncomeOutput, error) {
	var income entity.Income
	c.Repository.First(&income, input.ID)

	if income.ID == 0 {
		return UpdateIncomeOutput{}, errors.New("Income not found")
	}

	c.Repository.Model(&income).UpdateColumns(input.UpdatedIncome)

	output := UpdateIncomeOutput{
		Model: income.Model,
		Description: income.Description,
		Value: income.Value,
		Date: income.Date,
	}

	return output, nil


}
