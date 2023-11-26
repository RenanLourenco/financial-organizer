package month_resume

import (
	"fmt"
	expenseService "github.com/RenanLourenco/financial-organizer/expense/service"
	incomeService "github.com/RenanLourenco/financial-organizer/income/service"
	"gorm.io/gorm"
)

type MonthResume struct {
	Repository *gorm.DB
}

func (c *MonthResume) Execute(input MonthResumeInput) (MonthResumeOutput, error) {
	var incomesValue float64
	var expensesValue float64

	incomesValue = 0
	incomeService := incomeService.IncomeService{
		Repository: c.Repository,
	}
	
	expensesValue = 0
	expenseService := expenseService.ExpenseService{
		Repository: c.Repository,
	}


	incomes, err := incomeService.ListIncomesByMonth(input.Month,input.Year); if err != nil { fmt.Println(err.Error()) }

	for _, income := range incomes {
		incomesValue += income.Value
	}

	expenses, err := expenseService.ListByMonth(input.Month,input.Year); if err != nil { fmt.Println(err.Error()) }

	for _, expense := range expenses {
		expensesValue += expense.Value 
	}

	balance := incomesValue - expensesValue

	expensesByCategory, err := expenseService.ListExpenseEachCategoryByMonth(input.Month,input.Year); if err != nil {fmt.Println(err.Error())}

	

	output := MonthResumeOutput{
		ExpensesTotal: expensesValue,
		IncomesTotal: incomesValue,
		Balance: balance,
		TotalEachCategory: expensesByCategory,
	}

	return output, nil

}