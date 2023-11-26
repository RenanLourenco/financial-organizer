package entity

import (
	expense "github.com/RenanLourenco/financial-organizer/expense/adapter/entity"
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Description string `json:"description" validate:"nonzero"`
	Expenses []expense.Expense `gorm:"foreignKey:CategoryID"`
}

func ValidateIncome(category *Category) error{
	if err := validator.Validate(category); err != nil {
		return err
	}
	return nil
}