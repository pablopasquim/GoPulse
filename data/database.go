package data

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB // var global

func Connect() { // db config
	connection := "user=postgres dbname=stock password=piloto9815 host=localhost sslmode=disable"
	database, err := gorm.Open(postgres.Open(connection), &gorm.Config{}) // GORM open postgres connection
	if err != nil {
		log.Fatal("Error connection in database", err)
	}

	DB = database

}
