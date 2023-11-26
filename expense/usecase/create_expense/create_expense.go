package create_expense

import (
	"errors"
	"fmt"
	"time"
	category "github.com/RenanLourenco/financial-organizer/categories/service"
	"github.com/RenanLourenco/financial-organizer/expense/adapter/entity"
	"gorm.io/gorm"
)

type CreateExpense struct {
	Repository *gorm.DB
	CategoryService *category.CategoryService
}

func (c *CreateExpense) Execute(input CreateExpenseDtoInput) (CreateExpenseDtoOutput, error){
	var expense entity.Expense
	var find entity.Expense

	if input.Category == "" {
		input.Category = "Others"
	}

	category, err := c.CategoryService.Categorize(input.Category)

	if err != nil {
		fmt.Println(err)
	}

	expense.Description = input.Description
	expense.Value = input.Value
	expense.Date = input.Date
	expense.CategoryID = category.ID

	
	c.Repository.Where(&entity.Expense{Description: expense.Description}).First(&find)

	if find.ID != 0 {
		dateString := find.Date
		date, _ := time.Parse(time.RFC3339Nano, dateString)
		month := date.Month()

		inputDate, _ := time.Parse(time.RFC3339Nano, expense.Date)

		inputMonth := inputDate.Month()

		if month == inputMonth {
			return CreateExpenseDtoOutput{},errors.New("Expense already exists.")
		}

	}


	c.Repository.Create(&expense)

	output := CreateExpenseDtoOutput{}
	output.ID = expense.ID
	output.Category = category
	output.Description = expense.Description
	output.Date = expense.Date
	output.Value = expense.Value
	output.CreatedAt = expense.CreatedAt
	output.DeletedAt = expense.DeletedAt
	
	return output,nil


}