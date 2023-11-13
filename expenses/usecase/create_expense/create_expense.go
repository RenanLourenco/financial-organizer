package create_expense

import (
	"errors"
	"fmt"
	"time"

	expenses "github.com/RenanLourenco/financial-organizer/expenses/adapter/entity"
	"gorm.io/gorm"
)

type CreateExpense struct {
	Repository *gorm.DB
}

func (c *CreateExpense) Execute(input CreateExpenseDtoInput) (CreateExpenseDtoOutput, error){
	var expense expenses.Expenses
	var find expenses.Expenses

	expense.Description = input.Description
	expense.Value = input.Value
	expense.Date = input.Date

	
	c.Repository.Where(&expenses.Expenses{Description: expense.Description}).First(&find)

	fmt.Println(find.ID)

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
	output.Description = expense.Description
	output.Date = expense.Date
	output.Value = expense.Value
	output.CreatedAt = expense.CreatedAt
	output.DeletedAt = expense.DeletedAt
	
	return output,nil


}