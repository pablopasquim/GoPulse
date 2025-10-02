package data

import (
	"log"

	"github.com/pablopasquim/GoPulse/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB // var global

func Connect() { // db config
	connection := "host=localhost port=5432 user=pablo dbname=stock_db password=Piloto9815 sslmode=disable"
	database, err := gorm.Open(postgres.Open(connection)) // GORM open postgres connection
	if err != nil {
		log.Fatal("Error connection in database", err)
	}

	DB = database
	DB.AutoMigrate(&models.Item{})

}
