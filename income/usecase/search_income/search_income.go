package search_income

import (
	"github.com/RenanLourenco/financial-organizer/income/adapter/entity"
	"gorm.io/gorm"
)

type SearchIncome struct {
	Repository *gorm.DB
}

func (c *SearchIncome) Execute(input SearchIncomeInput) SearchIncomeOutput {
	var incomes []entity.Income

	searchString := "%" + input.Search + "%"

	c.Repository.Where("description LIKE ?", searchString).Find(&incomes)

	output := SearchIncomeOutput{
		Data: incomes,
	}

	return output

}
