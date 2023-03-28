package database

import (
	"fmt"
	"log"

	"challenge3/package/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "ilhamyd12"
	dbPort   = "8080"
	dbName   = "books"
)

func Init() *gorm.DB {
	dbURL := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbName, dbPort)

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.Debug().AutoMigrate(models.Book{})

	return db
}
