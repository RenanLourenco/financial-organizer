package expenses_routes

import (
	controller "github.com/RenanLourenco/financial-organizer/expenses/http"
	"github.com/gin-gonic/gin"
)


func SetupRoutes(r *gin.Engine){
	r.POST("/expenses",controller.CreateExpense)
	r.DELETE("/expenses/:id",controller.DeleteExpense)
	r.PUT("/expenses/:id",controller.UpdateExpense)
}