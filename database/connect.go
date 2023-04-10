package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/pratiksha/blogbackend/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	dsn := os.Getenv("DSN")
	database, error := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if error != nil {
		panic("could not connect to the db")
	} else {
		log.Print("connect succesfully to mysql db")
	}

	DB = database
	database.AutoMigrate(
		&models.User{},
		&models.Blog{},
	)

}
