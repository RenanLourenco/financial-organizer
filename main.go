package main

import (
	expense_routes "github.com/RenanLourenco/financial-organizer/expense/http/server"
	income_routes "github.com/RenanLourenco/financial-organizer/income/http/server"
	"github.com/RenanLourenco/financial-organizer/infra/database"
	"github.com/gin-gonic/gin"
)


func main(){
	database.ConnectDatabase()
	r := gin.Default()
	expense_routes.SetupRoutes(r)
	income_routes.SetupRoutes(r)

	r.Run(":5000")
}