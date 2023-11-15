package list_expenses

import (
	expenses "github.com/RenanLourenco/financial-organizer/expenses/adapter/entity"
)

type ListExpenseInput struct {
	ID string `json:"id"`
}

type ListExpenseOutput struct {
	Data []expenses.Expenses
}