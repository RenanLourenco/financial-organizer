package create_income

import (
	"errors"
	"time"

	"github.com/RenanLourenco/financial-organizer/income/adapter/entity"
	"gorm.io/gorm"
)

type CreateIncome struct {
	Repository *gorm.DB
}
// Execute Create a new income
// @Summary Create a new income
// @Description Create a new income
// @Produce json
// @Tags Income
// @Success 200 {object} CreateIncomeOutput
// @Router /incomes [post]
func (c *CreateIncome) Execute(input CreateIncomeInput) (CreateIncomeOutput, error) {
	var income entity.Income
	var find entity.Income

	income.Description = input.Description
	income.Date = input.Date
	income.Value = input.Value


	c.Repository.Where(&entity.Income{Description: income.Description}).First(&find)

	if find.ID != 0 {
		dateString := find.Date
		date, _ := time.Parse(time.RFC3339Nano, dateString)
		month := date.Month()

		inputDate, _ := time.Parse(time.RFC3339Nano, income.Date)

		inputMonth := inputDate.Month()

		if month == inputMonth {
			return CreateIncomeOutput{},errors.New("Expense already exists.")
		}
	}

	c.Repository.Create(&income)

	output := CreateIncomeOutput{}
	output.ID = income.ID
	output.Description = income.Description
	output.Date = income.Date
	output.Value = income.Value
	output.CreatedAt = income.CreatedAt
	output.DeletedAt = income.DeletedAt
	
	return output,nil

}