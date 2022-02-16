package main

import (
	"log"

	"github.com/carloshdurante/geolocation_api/api"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	api.Run()
}
