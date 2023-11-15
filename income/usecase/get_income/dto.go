package get_income

import "gorm.io/gorm"


type GetIncomeInput struct {
	ID string `json:"id"`
}

type GetIncomeOutput struct {
	gorm.Model
	Description string `json:"description"`
	Value      float64 `json:"value"`
	Date string `json:"date"`
}