package database

import (
	"fmt"
	"log"
	"os"
	category "github.com/RenanLourenco/financial-organizer/categories/adapter/entity"
	expense "github.com/RenanLourenco/financial-organizer/expense/adapter/entity"
	income "github.com/RenanLourenco/financial-organizer/income/adapter/entity"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbUser string
var dbPassword string
var dbName string
var dbHost string

var (
	DB *gorm.DB
	err error
)

func loadEnviromentVariables(){
	dbUser = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbName = os.Getenv("DB_NAME")
	dbHost = os.Getenv("DB_HOST")
}

func ConnectDatabase(){
	err := godotenv.Load(".env")
	if err != nil {
		error := godotenv.Load("../.env"); if error != nil { log.Fatalln(error.Error())}
	}

	loadEnviromentVariables()
	
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=America/Sao_Paulo",dbHost,dbUser,dbPassword,dbName)

	DB, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Database connected!")
	DB.AutoMigrate(&category.Category{},&expense.Expense{},&income.Income{})
}


