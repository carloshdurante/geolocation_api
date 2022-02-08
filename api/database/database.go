package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDb(){
	start := ("postgres://root:123456@localhost:5432/api_golang_db?sslmode=disable")

	database, err := gorm.Open(postgres.Open(start), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}

	db = database

	log.Println("Database is running")
}
