package main

import (
	"go-mock-postgre/database"
	"log"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	if err := database.AddUser("foo"); err != nil {
		log.Printf("Error: %v\n", err)
	}
}
