package repository

import (
	"context"
	"database/sql"
	"log"
	"os"
	"tangapp-be/config"
	"tangapp-be/repository"
	"tangapp-be/utils"
	"testing"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/lib/pq"
)

var testDB *pgxpool.Pool
var testQueries *repository.Queries

func NullString(ns string) sql.NullString {
	return utils.ToNullString(ns)
}

func NullTime(nt time.Time) sql.NullTime {
	return utils.ToNullTime(nt)
}

func NullInt32(i int32) sql.NullInt32 {
	return utils.ToNullInt32(i)
}

func TestMain(m *testing.M) {
	// Init DB
	config, err := config.LoadConfig("../")
	if err != nil {
		log.Fatal(err)
	}

	// Use DB connection string from the environment variable
	connString := config.DBCredential // DB_CREDENTIAL should be a PostgreSQL URI

	// Initialize a connection pool to the database
	testDB, err = pgxpool.Connect(context.Background(), connString)
	if err != nil {
		log.Fatalf("Could not connect to test database: %v", err)
	}
	defer testDB.Close()

	// Initialize `testQueries` with the `testDB` connection
	testQueries = repository.New(testDB)

	// Run tests
	code := m.Run()

	os.Exit(code)
}
