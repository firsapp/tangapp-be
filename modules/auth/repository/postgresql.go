package postgresql

import (
	"tangapp-be/repository"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Auth defines a postgres repository with the required dependencies.
type Auth struct {
	conn *pgxpool.Pool
	q    *repository.Queries
}

// NewAuth returns an instance of Todo repository.
func NewAuth(conn *pgxpool.Pool) *Auth {
	return &Auth{
		conn: conn,
		q:    repository.New(conn),
	}
}