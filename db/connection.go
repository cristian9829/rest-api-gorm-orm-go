package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DSN = "host=localhost user=cris password=pass dbname=gorm port=5432"
var DB *gorm.DB

func DBConnection() {
	//Todo
	var error error
	DB, error = gorm.Open(postgres.Open(DSN), &gorm.Config{})

	if error != nil {
		log.Fatal(error)
	} else {
		log.Println("DB CONNECTED")
	}
}
