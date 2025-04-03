package database

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func databaseConnection() (*gorm.DB, error) {
	databasePath := os.Getenv("DATABASE_PATH")
	database, err := gorm.Open(sqlite.Open(databasePath), &gorm.Config{})
	if err != nil {
		log.Printf("Failed to connect to database")
		return nil, err
	}
	return database, err
}
