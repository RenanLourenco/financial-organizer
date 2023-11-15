package list_expenses

import (
	"github.com/RenanLourenco/financial-organizer/expense/adapter/entity"
)

type ListExpenseInput struct {
	ID string `json:"id"`
}

type ListExpenseOutput struct {
	Data []entity.Expense
}