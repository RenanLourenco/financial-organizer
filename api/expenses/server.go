package expenses_routes

import (
	controller "github.com/RenanLourenco/financial-organizer/expenses/http"
	"github.com/gin-gonic/gin"
)


func SetupRoutes(r *gin.Engine){
	r.POST("/expenses",controller.CreateExpense)
}