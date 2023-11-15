package update_expense

import (
	expenses "github.com/RenanLourenco/financial-organizer/expenses/adapter/entity"
	"gorm.io/gorm"
)

type UpdateExpenseInput struct {
	ID string `json:"id"`
	UpdatedExpense expenses.Expenses
}

type UpdateExpenseOutput struct {
	gorm.Model
	Description string `json:"description"`
	Value float64 `json:"value"`
	Date string `json:"date"`
}