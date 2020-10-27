package database

import (
	"log"

	"github.com/passwdapp/box/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var dbConnection *gorm.DB

// GetDBConnection is used to globally access the gorm connection
func GetDBConnection() *gorm.DB {
	return dbConnection
}

// Connect is used to connect to the gorm db
func Connect(path string) {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.User{})

	dbConnection = db
}
