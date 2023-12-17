package main

import (
	expense_routes "github.com/RenanLourenco/financial-organizer/expense/http/server"
	income_routes "github.com/RenanLourenco/financial-organizer/income/http/server"
	"github.com/RenanLourenco/financial-organizer/infra/database"
	resume_routes "github.com/RenanLourenco/financial-organizer/resume/http/server"
	"github.com/gin-gonic/gin"

	_ "github.com/RenanLourenco/financial-organizer/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)



// @title Financial Organizer API
// @version 1.0.0
// @description App to organize all your financial life

// @contact.name Renan Lourenço
// @contact.url https://github.com/RenanLourenco
// @contact.email renanloub@protonmail.com

// @host localhost:5000

func main(){
	database.ConnectDatabase()
	r := gin.Default()
	gin.SetMode("release")
	expense_routes.SetupRoutes(r)
	income_routes.SetupRoutes(r)
	resume_routes.SetupRoutes(r)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":5000")
}