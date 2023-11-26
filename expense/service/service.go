package service

import (
	"errors"
	"fmt"
	"time"

	categoryEntity "github.com/RenanLourenco/financial-organizer/categories/adapter/entity"
	"github.com/RenanLourenco/financial-organizer/expense/adapter/entity"
	"gorm.io/gorm"
)

type ExpenseService struct {
	Repository *gorm.DB
}

func (c *ExpenseService) ListByCategory(category string) []entity.Expense {
	var expenses []entity.Expense

	var categoryFind categoryEntity.Category

	c.Repository.Where("description = ?", category).First(&categoryFind)

	if categoryFind.ID == 0 {
		return []entity.Expense{}
	}

	c.Repository.Where("category_id = ?", categoryFind.ID).Find(&expenses)

	return expenses

}

func (c *ExpenseService) ListByMonth(month string, year string) ([]entity.Expense, error) {
	var expenses []entity.Expense

	yearInt := 0
	monthInt := 0

	_, err := fmt.Sscanf(year, "%d", &yearInt)
	if err != nil {
		return expenses, errors.New("Invalid year or month format")
	}
	_, err = fmt.Sscanf(month, "%d", &monthInt)
	if err != nil {
		return expenses, errors.New("Invalid year or month format")
	}

	layout := "2006-01-02T15:04:05"
	startOfMonth := time.Date(yearInt, time.Month(monthInt), 1, 0, 0, 0, 0, time.UTC)
	endOfMonth := startOfMonth.AddDate(0, 1, 0).Add(-time.Nanosecond)

	c.Repository.Where("date BETWEEN ? AND ?", startOfMonth.Format(layout), endOfMonth.Format(layout)).Find(&expenses)

	return expenses, nil
}

func (c *ExpenseService) ListExpenseEachCategoryByMonth(month string, year string) ([]ExpensesByCategory, error){
	var expenseByCategory []ExpensesByCategory


	expenses, err := c.ListByMonth(month,year); if err != nil { return expenseByCategory, errors.New(err.Error()) }
	
	for _, expense := range expenses {
		var category categoryEntity.Category
		var categoryMap ExpensesByCategory

		c.Repository.First(&category, expense.CategoryID)

		categoryMap.Category = category.Description 
		categoryMap.Value += expense.Value

		expenseByCategory = append(expenseByCategory, categoryMap)

	}

	return expenseByCategory, nil


}
