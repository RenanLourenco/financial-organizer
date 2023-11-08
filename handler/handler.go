package handler

import (
	"github.com/gin-gonic/gin"
)

func handlerRequests(){
	r := gin.Default()
	r.POST("/expenses")
}