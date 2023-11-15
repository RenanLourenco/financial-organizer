package income_routes

import (
	controller "github.com/RenanLourenco/financial-organizer/income/http"
	"github.com/gin-gonic/gin"
)


func SetupRoutes(r *gin.Engine) {
	r.POST("/incomes", controller.CreateIncome)
}