package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/RenanLourenco/financial-organizer/income/adapter/entity"
	"gorm.io/gorm"
)

type IncomeService struct {
	Repository *gorm.DB
}

func (c *IncomeService) ListIncomesByMonth(month string, year string) ([]entity.Income, error) {
	var incomes []entity.Income

	yearInt := 0
	monthInt := 0

	_, err := fmt.Sscanf(year, "%d", &yearInt)
	if err != nil {
		return incomes, errors.New("Invalid year or month format")
	}
	_, err = fmt.Sscanf(month, "%d", &monthInt)
	if err != nil {
		return incomes, errors.New("Invalid year or month format")
	}

	layout := "2006-01-02T15:04:05"
	startOfMonth := time.Date(yearInt, time.Month(monthInt), 1, 0, 0, 0, 0, time.UTC)
	endOfMonth := startOfMonth.AddDate(0, 1, 0).Add(-time.Nanosecond)

	c.Repository.Where("date BETWEEN ? AND ?", startOfMonth.Format(layout), endOfMonth.Format(layout)).Find(&incomes)

	return incomes, nil
}
