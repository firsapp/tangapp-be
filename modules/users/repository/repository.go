package repository

import (
	"context"
	"database/sql"
	"tangapp-be/errors"
	"tangapp-be/queries"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4" // Harus pake v4. or else error "no rows" gak works
	"github.com/jackc/pgx/v4/pgxpool"
)

type UserRepository struct {
	q *queries.Queries
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		q: queries.New(db),
	}
}

// Adds a new user to the database
func (r *UserRepository) AddUser(ctx context.Context, username, email string) (queries.User, error) {
	arg := queries.AddUserParams{
		Username: sql.NullString{String: username, Valid: true},
		Email:    email,
	}

	user, err := r.q.AddUser(ctx, arg)
	if err != nil {
		return queries.User{}, err
	}
	return user, nil
}

// Gets a user by id (gak tau uuid.UUID bener atau enggak)
func (r *UserRepository) GetUserByID(ctx context.Context, ID uuid.UUID) (queries.User, error) {
	user, err := r.q.GetUserByID(ctx, ID)
	// TO-DO : Convert nullable payload into string
	if err != nil {
		if err == pgx.ErrNoRows {
			// User not found
			return queries.User{}, &errors.UserNotFoundError{ID: ID}
		}
		return queries.User{}, err
	}

	return user, nil
}
