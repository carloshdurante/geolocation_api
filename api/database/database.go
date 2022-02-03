package database

import (
	"log"
	"fmt"

	"github.com/carloshdurante/geolocation_api/api/controllers"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDb(server *Server){
	start := "postgres://username:password@localhost:5432/geolocation_api?sslmode=disable"

	database, err := gorm.Open(postgres.Open(start), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}

	db = database


	fmt.Println("Database connected.")
	log.Print("Database connected.")
}
