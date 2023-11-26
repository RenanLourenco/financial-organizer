package list_expense_by_month

import "github.com/RenanLourenco/financial-organizer/expense/adapter/entity"

type ListExpenseByMonthInput struct {
	Year string `json:"year"`
	Month string `json:"month"`
}
type ListExpenseByMonthOutput struct {
	Data []entity.Expense
}