package create_income

import "gorm.io/gorm"

type CreateIncomeInput struct {
	Description string `json:"description"`
	Value float64 `json:"value"`
	Date string `json:"date"`
}

type CreateIncomeOutput struct {
	gorm.Model
	Description string `json:"description"`
	Value      float64 `json:"value"`
	Date string `json:"date"`
}