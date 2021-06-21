package configs

import (
	"bookstore/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func ConnectDataBase() {
	database, err := gorm.Open("sqlite3", "test.db")
	book := models.Book{}

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&book)
	DB = database
}
