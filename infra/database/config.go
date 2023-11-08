package database

import (
	"fmt"
	"log"
	"os"

	expenses "github.com/RenanLourenco/financial-organizer/expenses/adapter/entity"
	income "github.com/RenanLourenco/financial-organizer/income/adapter/entity"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbUser string
var dbPassword string
var dbName string

var (
	DB *gorm.DB
	err error
)

func loadEnviromentVariables(){
	dbUser = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbName = os.Getenv("DB_NAME")
}

func ConnectDatabase(){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	loadEnviromentVariables()
	
	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=America/Sao_Paulo",dbUser,dbPassword,dbName)

	DB, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Database connected!")
	DB.AutoMigrate(&expenses.Expenses{},&income.Income{})
}


