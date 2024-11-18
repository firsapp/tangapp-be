package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testDB *sql.DB
var testQueries *Queries

func TestMain(m *testing.M) {
	// Connect to the test database
	var err error
	fmt.Println(os.Getenv("DB_CREDENTIAL"))
	testDB, err = sql.Open("postgres", "redacted")
	if err != nil {
		log.Fatalf("Could not connect to test database: %v", err)
	}
	defer testDB.Close()

	// Initialize `testQueries` with the `testDB` connection
	testQueries = New(testDB)

	// Run migrations (uncomment these if using goose for migrations)
	// if err := goose.Up(testDB, "../db/migration"); err != nil {
	// 	log.Fatalf("Could not run migrations: %v", err)
	// }

	// Run tests
	code := m.Run()

	// Teardown: rollback all migrations if needed
	// if err := goose.Down(testDB, "../db/migration"); err != nil {
	// 	log.Fatalf("Could not rollback migrations: %v", err)
	// }

	os.Exit(code)
}
