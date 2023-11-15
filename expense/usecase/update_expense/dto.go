package update_expense

import (
	"github.com/RenanLourenco/financial-organizer/expense/adapter/entity"
	"gorm.io/gorm"
)

type UpdateExpenseInput struct {
	ID string `json:"id"`
	UpdatedExpense entity.Expense
}

type UpdateExpenseOutput struct {
	gorm.Model
	Description string `json:"description"`
	Value float64 `json:"value"`
	Date string `json:"date"`
}