package list_income_by_month

import "github.com/RenanLourenco/financial-organizer/income/adapter/entity"

type ListIncomeByMonthInput struct {
	Year string `json:"year"`
	Month string `json:"month"`
}

type ListIncomeByMonthOutput struct {
	Data []entity.Income
}