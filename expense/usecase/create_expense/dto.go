package create_expense

import (
	"gorm.io/gorm"
)

type CreateExpenseDtoInput struct {
	Description string `json:"description"`
	Value float64 `json:"value"`
	Date string `json:"date"`
}

type CreateExpenseDtoOutput struct{
	gorm.Model
	Description string `json:"description"`
	Value      float64 `json:"value"`
	Date string `json:"date"`
}