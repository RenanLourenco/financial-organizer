package search_expense

import "github.com/RenanLourenco/financial-organizer/expense/adapter/entity"

type SearchExpenseInput struct {
	Search string 
}

type SearchExpenseOutput struct{
	Data []entity.Expense `json:"data"`
}
