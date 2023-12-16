package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	expense "github.com/RenanLourenco/financial-organizer/expense/adapter/entity"
	eController "github.com/RenanLourenco/financial-organizer/expense/http"
	eCreateUsecase "github.com/RenanLourenco/financial-organizer/expense/usecase/create_expense"
	eDeleteUsecase "github.com/RenanLourenco/financial-organizer/expense/usecase/delete_expense"
	eGetUsecase "github.com/RenanLourenco/financial-organizer/expense/usecase/get_expense"
	eListUsecase "github.com/RenanLourenco/financial-organizer/expense/usecase/list_expenses"
	eUpdateUsecase "github.com/RenanLourenco/financial-organizer/expense/usecase/update_expense"
	income "github.com/RenanLourenco/financial-organizer/income/adapter/entity"
	iController "github.com/RenanLourenco/financial-organizer/income/http"
	iCreateUsecase "github.com/RenanLourenco/financial-organizer/income/usecase/create_income"
	iDeleteUsecase "github.com/RenanLourenco/financial-organizer/income/usecase/delete_income"
	iGetUsecase "github.com/RenanLourenco/financial-organizer/income/usecase/get_income"
	iListUseCase "github.com/RenanLourenco/financial-organizer/income/usecase/list_incomes"
	iUpdateUsecase "github.com/RenanLourenco/financial-organizer/income/usecase/update_income"
	"github.com/RenanLourenco/financial-organizer/infra/database"
	rController "github.com/RenanLourenco/financial-organizer/resume/http"
	rMonthResume "github.com/RenanLourenco/financial-organizer/resume/usecase/month_resume"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

const layout = "2006-01-02T15:04:05.999Z"

func SetupRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	return r
}

func CleanDatabases() {
	sqlStringExpenses := fmt.Sprintf("DELETE from expenses")
	database.DB.Exec(sqlStringExpenses)
	sqlStringIncomes := fmt.Sprintf("DELETE from incomes")
	database.DB.Exec(sqlStringIncomes)
}

func CreateExpenseMock() expense.Expense {
	var createdExpense expense.Expense

	createdExpense.Description = "mock expense 1"
	createdExpense.Date = "2023-11-19T15:04:05.999Z"
	createdExpense.Value = 301
	createdExpense.CategoryID = 2

	database.DB.Create(&createdExpense)

	return createdExpense
}

func CreateIncomeMock() income.Income {
	var createdIncome income.Income

	createdIncome.Description = "mock income 1"
	createdIncome.Date = "2023-11-19T15:04:05.999Z"
	createdIncome.Value = 301

	database.DB.Create(&createdIncome)

	return createdIncome
}

func CreateIncomeForResumeMock() income.Income {
	var createdIncome income.Income

	createdIncome.Description = "Income for Resume"
	createdIncome.Date = "2023-05-30T15:04:05.999Z"
	createdIncome.Value = 5000

	database.DB.Create(&createdIncome)

	return createdIncome
}



func CreateExpensesForResumeMock() []expense.Expense {
	expensesMocks := [5]eCreateUsecase.CreateExpenseDtoInput{
		0: {
			Description: "mock expense 1",
			Value:       100,
			Date:        "2023-05-19T15:04:05.999Z",
		},
		1: {
			Description: "mock expense 2",
			Value:       100,
			Date:        "2023-05-20T15:04:05.999Z",
		},
		2: {
			Description: "mock expense 3",
			Value:       100,
			Date:        "2023-05-21T15:04:05.999Z",
		},
		3: {
			Description: "mock expense 4",
			Value:       100,
			Date:        "2023-05-22T15:04:05.999Z",
		},
		4: {
			Description: "mock expense 5",
			Value:       100,
			Date:        "2023-05-23T15:04:05.999Z",
		},
	}

	var createdExpensesMocks []expense.Expense

	for _, expenseMock := range expensesMocks {
		var createdExpense expense.Expense

		createdExpense.Description = expenseMock.Description
		createdExpense.Date = expenseMock.Date
		createdExpense.Value = expenseMock.Value
		createdExpense.CategoryID = 2

		database.DB.Create(&createdExpense)

		createdExpensesMocks = append(createdExpensesMocks, createdExpense)
	}

	return createdExpensesMocks
}

func CreateExpensesMocks() {
	expensesMocks := [5]eCreateUsecase.CreateExpenseDtoInput{
		0: {
			Description: "mock expense 1",
			Value:       301,
			Date:        "2023-11-19T15:04:05.999Z",
		},
		1: {
			Description: "mock expense 2",
			Value:       301,
			Date:        "2023-11-20T15:04:05.999Z",
		},
		2: {
			Description: "mock expense 3",
			Value:       301,
			Date:        "2023-11-21T15:04:05.999Z",
		},
		3: {
			Description: "mock expense 4",
			Value:       301,
			Date:        "2023-11-22T15:04:05.999Z",
		},
		4: {
			Description: "mock expense 5",
			Value:       301,
			Date:        "2023-11-23T15:04:05.999Z",
		},
	}

	var createdExpensesMocks []expense.Expense

	for _, expenseMock := range expensesMocks {
		var createdExpense expense.Expense

		createdExpense.Description = expenseMock.Description
		createdExpense.Date = expenseMock.Date
		createdExpense.Value = expenseMock.Value
		createdExpense.CategoryID = 2

		database.DB.Create(&createdExpense)

		createdExpensesMocks = append(createdExpensesMocks, createdExpense)
	}

}

func CreateIncomesMocks() {
	incomesMocks := [5]iCreateUsecase.CreateIncomeInput{
		0: {
			Description: "mock income 1",
			Value:       301,
			Date:        "2023-11-19T15:04:05.999Z",
		},
		1: {
			Description: "mock income 2",
			Value:       301,
			Date:        "2023-11-20T15:04:05.999Z",
		},
		2: {
			Description: "mock income 3",
			Value:       301,
			Date:        "2023-11-21T15:04:05.999Z",
		},
		3: {
			Description: "mock income 4",
			Value:       301,
			Date:        "2023-11-22T15:04:05.999Z",
		},
		4: {
			Description: "mock income 5",
			Value:       301,
			Date:        "2023-11-23T15:04:05.999Z",
		},
	}

	var createdIncomeMocks []income.Income

	for _, incomeMock := range incomesMocks {
		var createdIncome income.Income

		createdIncome.Description = incomeMock.Description
		createdIncome.Date = incomeMock.Date
		createdIncome.Value = incomeMock.Value

		database.DB.Create(&createdIncome)

		createdIncomeMocks = append(createdIncomeMocks, createdIncome)
	}

}

func CreateIncomesMonthMock(){
	incomesMocks := [5]iCreateUsecase.CreateIncomeInput{
		0: {
			Description: "mock income 1",
			Value:       301,
			Date:        "2023-06-19T15:04:05.999Z",
		},
		1: {
			Description: "mock income 2",
			Value:       301,
			Date:        "2023-06-20T15:04:05.999Z",
		},
		2: {
			Description: "mock income 3",
			Value:       301,
			Date:        "2023-06-21T15:04:05.999Z",
		},
		3: {
			Description: "mock income 4",
			Value:       301,
			Date:        "2023-06-22T15:04:05.999Z",
		},
		4: {
			Description: "mock income 5",
			Value:       301,
			Date:        "2023-06-23T15:04:05.999Z",
		},
	}

	var createdIncomeMocks []income.Income

	for _, incomeMock := range incomesMocks {
		var createdIncome income.Income

		createdIncome.Description = incomeMock.Description
		createdIncome.Date = incomeMock.Date
		createdIncome.Value = incomeMock.Value

		database.DB.Create(&createdIncome)

		createdIncomeMocks = append(createdIncomeMocks, createdIncome)
	}
}

func TestCreateExpense(t *testing.T) {
	database.ConnectDatabase()
	r := SetupRoutes()
	defer CleanDatabases()

	r.POST("/expenses", eController.CreateExpense)

	newExpense := eCreateUsecase.CreateExpenseDtoInput{
		Description: "Testing",
		Value:       300,
		Date:        "2023-12-08T11:04:03.279Z",
	}

	jsonValue, _ := json.Marshal(newExpense)

	req, _ := http.NewRequest("POST", "/expenses", bytes.NewBuffer(jsonValue))
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	var createdExpense eCreateUsecase.CreateExpenseDtoOutput
	json.Unmarshal(resp.Body.Bytes(), &createdExpense)

	assert.Equal(t, newExpense.Value, createdExpense.Value)
	assert.Equal(t, newExpense.Description, createdExpense.Description)
	assert.Equal(t, newExpense.Date, createdExpense.Date)
}

func TestGetExpense(t *testing.T) {
	database.ConnectDatabase()
	expenseMock := CreateExpenseMock()
	defer CleanDatabases()

	id := expenseMock.ID

	r := SetupRoutes()
	r.GET("/expenses/:id", eController.GetExpense)

	url := fmt.Sprintf("/expenses/%d", id)

	req, _ := http.NewRequest("GET", url, nil)

	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	var findExpense eGetUsecase.GetExpenseOutput

	json.Unmarshal(resp.Body.Bytes(), &findExpense)

	assert.Equal(t, findExpense.ID, id)

}

func TestListExpenses(t *testing.T) {
	database.ConnectDatabase()
	CreateExpensesMocks()
	defer CleanDatabases()
	r := SetupRoutes()
	r.GET("/expenses", eController.ListExpenses)
	req, _ := http.NewRequest("GET", "/expenses", nil)
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	var listExpenses eListUsecase.ListExpenseOutput

	json.Unmarshal(resp.Body.Bytes(), &listExpenses.Data)

	assert.True(t, len(listExpenses.Data) > 0)

}

func TestUpdateExpense(t *testing.T) {
	database.ConnectDatabase()
	expense := CreateExpenseMock()
	defer CleanDatabases()

	url := fmt.Sprintf("/expenses/%d", expense.ID)

	r := SetupRoutes()
	r.PUT("/expenses/:id", eController.UpdateExpense)

	expense.Description = "Updated Description"
	expense.Value = float64(3000)
	expense.Date = "2023-05-30T00:04:05.999Z"

	jsonExpense, _ := json.Marshal(expense)

	req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(jsonExpense))
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	var updatedExpense eUpdateUsecase.UpdateExpenseOutput

	json.Unmarshal(resp.Body.Bytes(), &updatedExpense)

	assert.Equal(t, updatedExpense.Description, "Updated Description")
	assert.Equal(t, updatedExpense.Value, float64(3000))
	assert.Equal(t, updatedExpense.Date, "2023-05-30T00:04:05.999Z")

}

func TestDeleteExpense(t *testing.T) {
	database.ConnectDatabase()
	expense := CreateExpenseMock()
	defer CleanDatabases()

	url := fmt.Sprintf("/expenses/%d", expense.ID)

	r := SetupRoutes()
	r.DELETE("/expenses/:id", eController.DeleteExpense)

	req, _ := http.NewRequest("DELETE", url, nil)
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	var response eDeleteUsecase.DeleteExpenseOutput

	json.Unmarshal(resp.Body.Bytes(), &response)

	assert.Equal(t, response.Msg, "Expense deleted.")

}

func TestCreateIncome(t *testing.T) {
	database.ConnectDatabase()
	r := SetupRoutes()
	defer CleanDatabases()

	r.POST("/incomes", iController.CreateIncome)

	newIncome := iCreateUsecase.CreateIncomeInput{
		Description: "new income",
		Value:       300,
		Date:        "2023-11-23T15:04:05.999Z",
	}

	jsonIncome, _ := json.Marshal(newIncome)

	req, _ := http.NewRequest("POST", "/incomes", bytes.NewBuffer(jsonIncome))
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	var response iCreateUsecase.CreateIncomeOutput

	json.Unmarshal(resp.Body.Bytes(), &response)

	assert.Equal(t, http.StatusCreated, resp.Code)

}

func TestGetIncome(t *testing.T) {
	database.ConnectDatabase()
	r := SetupRoutes()
	defer CleanDatabases()

	incomeMock := CreateIncomeMock()

	r.GET("/incomes/:id", iController.GetIncome)

	url := fmt.Sprintf("/incomes/%d", incomeMock.ID)

	req, _ := http.NewRequest("GET", url, nil)
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	var response iGetUsecase.GetIncomeOutput

	json.Unmarshal(resp.Body.Bytes(), &response)

	assert.Equal(t, response.ID, incomeMock.ID)
	assert.Equal(t, response.Description, incomeMock.Description)
	assert.Equal(t, response.Date, incomeMock.Date)

}

func TestListIncomes(t *testing.T) {
	database.ConnectDatabase()
	r := SetupRoutes()
	CreateIncomesMocks()
	defer CleanDatabases()

	r.GET("/incomes", iController.ListIncomes)
	req, _ := http.NewRequest("GET", "/incomes", nil)
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	var response iListUseCase.ListIncomesOutput

	json.Unmarshal(resp.Body.Bytes(), &response.Data)

	assert.True(t, len(response.Data) > 0)

}



func TestUpdateIncome(t *testing.T) {
	database.ConnectDatabase()
	r := SetupRoutes()
	incomeMock := CreateIncomeMock()
	defer CleanDatabases()

	r.PUT("/incomes/:id", iController.UpdateIncome)

	url := fmt.Sprintf("/incomes/%d", incomeMock.ID)

	incomeMock.Description = "Updated Description."
	incomeMock.Date = "2023-11-19T15:04:05.999Z"
	incomeMock.Value = float64(533)

	jsonIncome, _ := json.Marshal(incomeMock)

	req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(jsonIncome))
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	var response iUpdateUsecase.UpdateIncomeOutput

	json.Unmarshal(resp.Body.Bytes(), &response)

	assert.Equal(t, response.Description, "Updated Description.")
	assert.Equal(t, response.Date, "2023-11-19T15:04:05.999Z")
	assert.Equal(t, response.Value, float64(533))
}

func TestDeleteIncome(t *testing.T){
	database.ConnectDatabase()
	income := CreateIncomeMock()
	r := SetupRoutes()
	defer CleanDatabases()

	id := income.ID

	url := fmt.Sprintf("/incomes/%d", id)

	r.DELETE("/incomes/:id",iController.DeleteIncome)

	req,_ := http.NewRequest("DELETE", url, nil)
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp,req)

	var response iDeleteUsecase.DeleteIncomeOutput

	json.Unmarshal(resp.Body.Bytes(), &response)

	assert.Equal(t, response.Msg, "Income Deleted.")
}

func TestMonthResume(t *testing.T){
	database.ConnectDatabase()
	//Expense total 500
	createdExpenses := CreateExpensesForResumeMock()
	// Income total 5000
	createdIncomes := CreateIncomeForResumeMock()

	var finalBalance float64
	var allExpensesValue float64
	var allIncomesValue float64

	for _, expense := range createdExpenses {
		allExpensesValue += expense.Value
	}

	allIncomesValue = createdIncomes.Value

	finalBalance = allIncomesValue - allExpensesValue


	r := SetupRoutes()
	defer CleanDatabases()
	
	url := fmt.Sprintf("/resume/%s/%s","2023","05")

	r.GET("/resume/:year/:month",rController.Resume)

	
	req,_ := http.NewRequest("GET", url, nil)
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp,req)

	var response rMonthResume.MonthResumeOutput

	json.Unmarshal(resp.Body.Bytes(), &response)


	assert.Equal(t, response.Balance, finalBalance)
	assert.Equal(t, response.ExpensesTotal, allExpensesValue)
	assert.Equal(t, response.IncomesTotal, allIncomesValue)
	assert.True(t, len(response.TotalEachCategory) > 0)
	

}