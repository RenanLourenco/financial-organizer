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

func ValidateExpense(expense *Expenses) error{
	if err := validator.Validate(expense); err != nil {
		return err
	}
	return nil
}