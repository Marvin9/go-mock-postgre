package database

import (
	"github.com/jinzhu/gorm"
)

// Users -
type Users struct {
	gorm.Model
	Username string
}

// AddUser -
func AddUser(username string) error {
	db, err := ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	db.AutoMigrate(&Users{})

	err = db.Create(&Users{
		Username: username,
	}).Error

	if err != nil {
		return err
	}

	return nil
}
