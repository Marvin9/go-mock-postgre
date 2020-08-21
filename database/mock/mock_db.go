package mock

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

const (
	// CREATE - operation to create database
	CREATE = "CREATE"
	// DROP - operation to drop database
	DROP = "DROP"
)

// MockedDB is used in unit tests to mock db
func MockedDB(operation string) {
	if CI := os.Getenv("CI"); CI == "" {
		_, fileName, _, _ := runtime.Caller(0)
		currPath := filepath.Dir(fileName)
		err := godotenv.Load(currPath + "/../../.env")
		if err != nil {
			log.Fatalf("Error loading env.\n%v", err)
		}
	}

	dbName := os.Getenv("DATABASE_NAME")
	pgUser := os.Getenv("PSQL_USER")
	pgPassword := os.Getenv("PSQL_PASSWORD")

	var command string

	if operation == CREATE {
		command = "createdb"
	} else {
		command = "dropdb"
	}

	cmd := exec.Command(command, "-h", "localhost", "-U", pgUser, "-e", dbName)
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, fmt.Sprintf("PGPASSWORD=%v", pgPassword))

	if err := cmd.Run(); err != nil {
		log.Fatalf("Error executing %v on %v.\n%v", command, dbName, err)
	}
}
