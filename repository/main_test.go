package repository

import (
	"database/sql"
	"log"
	"os"
	"tangapp-be/config"
	"testing"

	_ "github.com/lib/pq"
)

var testDB *sql.DB
var testQueries *Queries

func TestMain(m *testing.M) {
	// Init DB
	config, err := config.LoadConfig("../")
	if err != nil {
		log.Fatal(err)
	}

	testDB, err = sql.Open("postgres", config.DBCredential)
	if err != nil {
		log.Fatalf("Could not connect to test database: %v", err)
	}
	defer testDB.Close()

	// Initialize `testQueries` with the `testDB` connection
	testQueries = New(testDB)

	// Run tests
	code := m.Run()

	os.Exit(code)
}
