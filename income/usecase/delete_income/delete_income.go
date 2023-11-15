package delete_income

import (
	"github.com/RenanLourenco/financial-organizer/income/adapter/entity"
	"gorm.io/gorm"
)

type DeleteIncome struct {
	Repository *gorm.DB
}

func (c *DeleteIncome) Execute(input DeleteIncomeInput) (DeleteIncomeOutput, error) {
	var income entity.Income

	c.Repository.Delete(&income, input.ID)

	output := DeleteIncomeOutput{
		Msg: "Income Deleted.",
	}

	return output, nil

}
