package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/pratiksha/blogbackend/database"
	"github.com/pratiksha/blogbackend/routes"
)

func main() {
	database.Connect()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading in env file")
	}
	port := os.Getenv("PORT")
	app := fiber.New()
	routes.SetUp(app)
	app.Listen(":" + port)
}
