package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/Nitesh-04/locked-in/config"
	"github.com/Nitesh-04/locked-in/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.ConnectDatabase()

	app := fiber.New()

	routes.UserRoutes(app)

	fmt.Println("Server running on port 8080")
	log.Fatal(app.Listen(":8080"))
}