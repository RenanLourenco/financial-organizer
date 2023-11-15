package update_income

import (
	"github.com/RenanLourenco/financial-organizer/income/adapter/entity"
	"gorm.io/gorm"
)


type UpdateIncomeInput struct {
	ID string `json:"id"`
	UpdatedIncome entity.Income
}

type UpdateIncomeOutput struct {
	gorm.Model
	Description string `json:"description"`
	Value float64 `json:"value"`
	Date string `json:"date"`
}