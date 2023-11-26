package create_expense

import (
	category "github.com/RenanLourenco/financial-organizer/categories/adapter/entity"
	"gorm.io/gorm"
)

type CreateExpenseDtoInput struct {
	Description string  `json:"description"`
	Value       float64 `json:"value"`
	Date        string  `json:"date"`
	Category    string  `json:"category"`
}

type CreateExpenseDtoOutput struct {
	gorm.Model
	Description string  `json:"description"`
	Value       float64 `json:"value"`
	Date        string  `json:"date"`
	Category    category.Category
}
