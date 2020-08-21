package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
)

// ConnectDB -
func ConnectDB() (*gorm.DB, error) {
	databaseURL := os.Getenv("DATABASE_URL")
	databaseName := os.Getenv("DATABASE_NAME")

	db, err := gorm.Open("postgres", fmt.Sprintf("%v/%v", databaseURL, databaseName))
	if err != nil {
		log.Printf("Error connecting database.\n%v", err)
		return nil, err
	}

	return db, nil
}
