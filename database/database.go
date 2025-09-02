package database

import (
	"log"

	"github.com/glebarez/sqlite"
	"github.com/voduybaokhanh/go-rest-api-starter/models"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// Use pure Go SQLite driver
	database, err := gorm.Open(sqlite.Open("books.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}

	database.AutoMigrate(&models.Book{})

	DB = database
}
