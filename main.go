package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/Nitesh-04/locked-in/config"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.ConnectDatabase()
	fmt.Println("Server is running...")
}