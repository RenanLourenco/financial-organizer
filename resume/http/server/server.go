package resume_routes

import (
	controller "github.com/RenanLourenco/financial-organizer/resume/http"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine){
	r.GET("/resume/:year/:month", controller.Resume)
}

