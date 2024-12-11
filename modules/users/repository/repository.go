package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"tangapp-be/errorx"
	"tangapp-be/queries"

	"github.com/google/uuid"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4" // Harus pake v4. or else error "no rows" gak works
	"github.com/jackc/pgx/v4/pgxpool"
)

type UserRepository struct {
	q *queries.Queries
}

var pgErr *pgconn.PgError

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
			return queries.User{}, &errorx.UserNotFoundError{ID: ID}
		}
		return queries.User{}, err
	}

	return user, nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, arg queries.UpdateUserParams) (queries.User, error) {
	user, err := r.q.UpdateUser(ctx, arg)
	if err != nil {
		fmt.Println(err)
		// Check if the error came from PostgreSQL-specific (pgconn.PgError)
		if errors.As(err, &pgErr) {
			return queries.User{}, &errorx.DatabaseError{
				Err: fmt.Errorf("PostgreSQL error: %s (Code: %s, Detail: %s)", pgErr.Message, pgErr.Code, pgErr.Detail),
			}
		}

		return queries.User{}, fmt.Errorf("unexpected error: %w", err)
	}
	return user, nil
}
