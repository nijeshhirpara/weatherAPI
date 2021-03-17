package main

import (
	"log"
	"weatherapi/webserver"

	"github.com/joho/godotenv"
)

func main() {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	log.Println("Trying to start server")

	webErr := webserver.Start()
	if webErr != nil {
		log.Fatal("Server cannot be able to start", webErr)
	}
}
