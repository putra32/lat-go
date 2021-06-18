package config

import (
	"bookstore/models"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func ConnectDataBase() {
	database, err := gorm.Open("go-sqlite3", "test.db")
	book := models.Book{}

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&book)
	DB = database
}
