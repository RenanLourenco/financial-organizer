package month_resume

import "github.com/RenanLourenco/financial-organizer/expense/service"


type MonthResumeInput struct {
	Month string
	Year string
}

type TotalCategoryExpense struct {
	Category string `json:"category"`
	Value float64 `json:"value"`
}

type MonthResumeOutput struct {
	ExpensesTotal float64 `json:"total_expenses"`
	IncomesTotal float64 `json:"total_incomes"`
	Balance float64 `json:"balance"`
	TotalEachCategory []service.ExpensesByCategory `json:"total_each_category"`
}