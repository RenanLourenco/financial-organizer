package entities

import (
	"time"

	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Expenses struct {
	gorm.Model
	Description string `json:"description" validate:"nonzero"`
	Value float64 `json:"value" validate:"nonzero"`
	Date time.Time `json:"date"`
}

func ValidateExpense(income *Income) error{
	if err := validator.Validate(income); err != nil {
		return err
	}
	return nil
}