package main

import (
	expenses_routes "github.com/RenanLourenco/financial-organizer/api/expenses"
	"github.com/RenanLourenco/financial-organizer/infra/database"
	"github.com/gin-gonic/gin"
)


func main(){
	database.ConnectDatabase()
	r := gin.Default()
	expenses_routes.SetupRoutes(r)

	r.Run(":5000")
}