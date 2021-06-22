package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

// DB Config represents db configuration
type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func BuildDBConfig() *DBConfig {
	db := DBConfig {
		Host: "localhost",
		Port: 3306,
		User: "root",
		Password: "",
		DBName: "first_go",
	}
	return &db
}

func DbURL(db *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?chartset=utf8&parseTime=True&loc=Local",
		db.User, db.Password, db.Host, db.Port, db.DBName,
	)
}