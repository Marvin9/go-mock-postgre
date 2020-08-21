package database_test

import (
	"go-mock-postgre/database"
	"go-mock-postgre/database/mock"
	"testing"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func TestAddUsers(t *testing.T) {
	mock.MockedDB(mock.CREATE)
	defer mock.MockedDB(mock.DROP)

	newUser := "fooo"
	database.AddUser(newUser)

	db, err := database.ConnectDB()
	if err != nil {
		t.Errorf("Error connecting database in %v\n%v", t.Name(), err)
	}
	defer db.Close()

	var newUserInDB database.Users
	notFound := db.Where("username = ?", newUser).First(&newUserInDB).RecordNotFound()

	if notFound {
		t.Errorf("%v was added but was not retrieved.", newUser)
	}
}
