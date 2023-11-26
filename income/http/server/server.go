package income_routes

import (
	controller "github.com/RenanLourenco/financial-organizer/income/http"
	"github.com/gin-gonic/gin"
)


func SetupRoutes(r *gin.Engine) {
	r.POST("/incomes", controller.CreateIncome)
	r.DELETE("/income/:id",controller.DeleteIncome)
	r.GET("/income/:id",controller.GetIncome)
	r.GET("/incomes",controller.ListIncomes)
	r.PUT("/income/:id",controller.UpdateIncome)
	r.GET("/incomes/:year/:month",controller.ListIncomeByMonth)
}