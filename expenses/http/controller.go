package controller

import (
	"net/http"

	expenses "github.com/RenanLourenco/financial-organizer/expenses/adapter/entity"
	"github.com/RenanLourenco/financial-organizer/expenses/usecase/create_expense"
	"github.com/RenanLourenco/financial-organizer/expenses/usecase/delete_expense"
	"github.com/RenanLourenco/financial-organizer/expenses/usecase/update_expense"
	"github.com/RenanLourenco/financial-organizer/infra/database"
	"github.com/gin-gonic/gin"
)


func CreateExpense(c *gin.Context){
	usecase := create_expense.CreateExpense{
		Repository: database.DB,
	}
	var input create_expense.CreateExpenseDtoInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
		})
		return
	}

	output, err := usecase.Execute(input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated,output)
}

func DeleteExpense(c *gin.Context){
	usecase := delete_expense.DeleteExpense{
		Repository: database.DB,
	}

	id := c.Params.ByName("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":"Id missing in query params.",
		})
		return
	}

	input := delete_expense.DeleteExpenseInput{
		ID: id,
	}

	output, err := usecase.Execute(input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK,output)
}

func UpdateExpense(c *gin.Context){
	usecase := update_expense.UpdateExpense{
		Repository: database.DB,
	}

	var expense expenses.Expenses

	if err := c.ShouldBindJSON(&expense); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":err.Error(),
		})
		return
	}

	id := c.Params.ByName("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":"Id missing in query params.",
		})
		return
	}

	input := update_expense.UpdateExpenseInput{
		ID: id,
		UpdatedExpense: expense,
	}

	output, err := usecase.Execute(input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK,output)

}