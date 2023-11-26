package controller

import (
	"net/http"

	"github.com/RenanLourenco/financial-organizer/income/adapter/entity"
	"github.com/RenanLourenco/financial-organizer/income/usecase/create_income"
	"github.com/RenanLourenco/financial-organizer/income/usecase/delete_income"
	"github.com/RenanLourenco/financial-organizer/income/usecase/get_income"
	"github.com/RenanLourenco/financial-organizer/income/usecase/list_income_by_month"
	"github.com/RenanLourenco/financial-organizer/income/usecase/list_incomes"
	"github.com/RenanLourenco/financial-organizer/income/usecase/search_income"
	"github.com/RenanLourenco/financial-organizer/income/usecase/update_income"
	"github.com/RenanLourenco/financial-organizer/infra/database"
	"github.com/gin-gonic/gin"
)

func CreateIncome(c *gin.Context) {
	usecase := create_income.CreateIncome{
		Repository: database.DB,
	}

	var input create_income.CreateIncomeInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	output, err := usecase.Execute(input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, output)

}

func DeleteIncome(c *gin.Context) {
	usecase := delete_income.DeleteIncome{
		Repository: database.DB,
	}

	id := c.Params.ByName("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Id missing in query params.",
		})
		return
	}

	input := delete_income.DeleteIncomeInput{
		ID: id,
	}

	output, err := usecase.Execute(input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, output)
}


func GetIncome(c *gin.Context) {
	usecase := get_income.GetIncome {
		Repository: database.DB,
	}

	id := c.Params.ByName("id")

	if id == "" {
		c.JSON( http.StatusBadRequest, gin.H{
			"msg": "Error missing ID.",
		})
		return
	}

	input := get_income.GetIncomeInput{
		ID:id,
	}

	output, err := usecase.Execute(input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, output)

}

func ListIncomes(c *gin.Context) {
	usecase := list_incomes.ListIncomes{
		Repository: database.DB,
	}

	queryDescription := c.Query("description")

	if queryDescription != ""{
		searchIncomeUsecase := search_income.SearchIncome{
			Repository: database.DB,
		}
		searchIncomeInput := search_income.SearchIncomeInput{
			Search: queryDescription,
		}

		output := searchIncomeUsecase.Execute(searchIncomeInput)

		c.JSON(http.StatusOK, output.Data)
		return
	}

	output, err := usecase.Execute()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, output.Data)
}

func UpdateIncome(c *gin.Context){
	usecase := update_income.UpdateIncome{
		Repository: database.DB,
	}

	var income entity.Income

	if err := c.ShouldBindJSON(&income); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	id := c.Params.ByName("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Id missing in query params.",
		})
		return
	}

	input := update_income.UpdateIncomeInput{
		ID:             id,
		UpdatedIncome: income,
	}

	output, err := usecase.Execute(input)


	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	
	c.JSON(http.StatusOK, output)

}

func ListIncomeByMonth(c *gin.Context){
	usecase := list_income_by_month.ListIncomeByMonth{
		Repository: database.DB,
	} 

	month := c.Params.ByName("month")
	year := c.Params.ByName("year")

	if year == "" || month == "" {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Year and month parameters are required"})
		return
	}

	input := list_income_by_month.ListIncomeByMonthInput{
		Month: month,
		Year:  year,
	}

	output, err := usecase.Execute(input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, output.Data)


}