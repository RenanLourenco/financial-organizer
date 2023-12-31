package list_income_by_month

import (
	"errors"
	"fmt"
	"time"

	"github.com/RenanLourenco/financial-organizer/income/adapter/entity"
	"gorm.io/gorm"
)

type ListIncomeByMonth struct {
	Repository *gorm.DB
}

// Execute Get incomes by month
// @Summary Get specific incomes by month
// @Description Get specific incomes by month
// @Tags Income
// @Success 200 {object} ListIncomeByMonthOutput
// @Router /incomes/:year/:month [get]
func (c *ListIncomeByMonth) Execute(input ListIncomeByMonthInput) (ListIncomeByMonthOutput, error){
	var incomes []entity.Income

	year := input.Year
	month := input.Month

	yearInt := 0
	monthInt := 0 

	_, err := fmt.Sscanf(year, "%d", &yearInt) ; if err != nil { return ListIncomeByMonthOutput{}, errors.New("Invalid year or month format")}
	_, err = fmt.Sscanf(month, "%d", &monthInt); if err != nil { return ListIncomeByMonthOutput{}, errors.New("Invalid year or month format")}

	layout := "2006-01-02T15:04:05"
	startOfMonth := time.Date(yearInt, time.Month(monthInt), 1, 0, 0, 0, 0, time.UTC)
	endOfMonth := startOfMonth.AddDate(0, 1, 0).Add(-time.Nanosecond)

	c.Repository.Where("date BETWEEN ? AND ?", startOfMonth.Format(layout), endOfMonth.Format(layout)).Find(&incomes)

	output := ListIncomeByMonthOutput{
		Data: incomes,
	}

	return output, nil

}