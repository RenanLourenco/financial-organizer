package expenses

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Expenses struct {
	gorm.Model
	Description string `json:"description" validate:"nonzero"`
	Value float64 `json:"value" validate:"nonzero"`
	Date string `json:"date"`
}

func (e *Expenses) ValidateExpense() error{
	if err := validator.Validate(e); err != nil {
		return err
	}
	return nil
}