package database

import (
	"os"

	"github.com/golang-crud/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
    err := godotenv.Load()
    if err != nil {
        panic("Error loading .env file")
    }

    dbUsername := os.Getenv("DB_USERNAME")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbName := os.Getenv("DB_NAME")

    dsn := dbUsername + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?parseTime=true"
    
    database, err := gorm.Open(mysql.Open(dsn))
    if err != nil {
        panic(err)
    }

   
   database.AutoMigrate(&models.Product{}) 

    DB = database
}
