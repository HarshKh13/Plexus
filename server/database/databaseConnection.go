package database

import (
	"log"
	"os"
	"server/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func DatabaseConnection() (*gorm.DB, error) {
	databasePath := os.Getenv("DATABASE_PATH")
	database, err := gorm.Open(sqlite.Open(databasePath), &gorm.Config{})
	if err != nil {
		log.Printf("Failed to connect to database")
		return nil, err
	}

	err = database.AutoMigrate(&models.Users{})
	return database, err
}
