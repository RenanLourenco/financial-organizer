package list_expense_by_month

import (
	"errors"
	"fmt"
	"time"
	"github.com/RenanLourenco/financial-organizer/expense/adapter/entity"
	"gorm.io/gorm"
)

type ListExpenseByMonth struct {
	Repository *gorm.DB
}


// Execute List expenses by month
// @Summary List expenses by month
// @Description List expenses by month
// @Produce json
// @Tags Expense
// @Success 200 {object} ListExpenseByMonthOutput
// @Router /expenses/:year/:month [get]
func (c *ListExpenseByMonth) Execute(input ListExpenseByMonthInput) (ListExpenseByMonthOutput, error) {
	var expenses []entity.Expense

	year := input.Year
	month := input.Month

	yearInt := 0
	monthInt := 0 

	_, err := fmt.Sscanf(year, "%d", &yearInt) ; if err != nil { return ListExpenseByMonthOutput{}, errors.New("Invalid year or month format")}
	_, err = fmt.Sscanf(month, "%d", &monthInt); if err != nil { return ListExpenseByMonthOutput{}, errors.New("Invalid year or month format")}

	layout := "2006-01-02T15:04:05"
	startOfMonth := time.Date(yearInt, time.Month(monthInt), 1, 0, 0, 0, 0, time.UTC)
	endOfMonth := startOfMonth.AddDate(0, 1, 0).Add(-time.Nanosecond)

	c.Repository.Where("date BETWEEN ? AND ?", startOfMonth.Format(layout), endOfMonth.Format(layout)).Find(&expenses)

	output := ListExpenseByMonthOutput{
		Data: expenses,
	}

	return output, nil

}