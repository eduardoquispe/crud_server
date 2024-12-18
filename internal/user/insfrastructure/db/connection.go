package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB;

func DbConnection() {
	var error error
	DB, error = gorm.Open(postgres.Open("host=localhost user=postgres password=123456 dbname=crud_db port=5432 sslmode=disable"))

	if error != nil {
		log.Fatal(error)
	} else {
		log.Println("Connection to database established successfully")
	}
}