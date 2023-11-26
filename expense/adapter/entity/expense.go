package entity

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Expense struct {
	gorm.Model
	Description string `json:"description" validate:"nonzero"`
	Value float64 `json:"value" validate:"nonzero"`
	Date string `json:"date"`
	CategoryID uint `json:"categoryID"`
}

func (e *Expense) ValidateExpense() error{
	if err := validator.Validate(e); err != nil {
		return err
	}
	return nil
}