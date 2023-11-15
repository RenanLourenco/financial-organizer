package controller

import (
	"net/http"

	"github.com/RenanLourenco/financial-organizer/income/usecase/create_income"
	"github.com/RenanLourenco/financial-organizer/infra/database"
	"github.com/gin-gonic/gin"
)

func CreateIncome(c *gin.Context){
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