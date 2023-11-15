package get_expense

import "gorm.io/gorm"

type GetExpenseInput struct {
	ID string `json:"id"`
}

type GetExpenseOutput struct {
	gorm.Model
	Description string `json:"description"`
	Value      float64 `json:"value"`
	Date string `json:"date"`
}