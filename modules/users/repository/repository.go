package repository

import (
	"context"
	"errors"
	"fmt"
	"tangapp-be/errorx"
	"tangapp-be/queries"
	"tangapp-be/utils"

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

type UserPayload struct {
	ID       string
	Username string
	Email    string
}

// Adds a new user to the database
func (r *UserRepository) AddUser(ctx context.Context, arg UserPayload) (queries.User, error) {

	user, err := r.q.AddUser(ctx, queries.AddUserParams{
		Username: utils.ToNullString(arg.Username),
		Email:    arg.Email,
	})
	if err != nil {
		return queries.User{}, err
	}

	return user, nil
}

// Gets a user by id (gak tau uuid.UUID bener atau enggak)
func (r *UserRepository) GetUserByID(ctx context.Context, ID string) (queries.User, error) {
	id := uuid.MustParse(ID)
	user, err := r.q.GetUserByID(ctx, id)
	// TO-DO : Convert nullable payload into string
	if err != nil {
		if err == pgx.ErrNoRows {
			// User not found
			return queries.User{}, &errorx.UserNotFoundError{ID: id}
		}
		return queries.User{}, err
	}

	return user, nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, arg UserPayload) (string, error) {
	user, err := r.q.UpdateUser(ctx, queries.UpdateUserParams{
		ID:       uuid.MustParse(arg.ID),
		Username: utils.ToNullString(arg.Username),
	})
	if err != nil {
		fmt.Println(err)
		// Check if the error came from PostgreSQL-specific (pgconn.PgError)
		if errors.As(err, &pgErr) {
			return "", &errorx.DatabaseError{
				Err: fmt.Errorf("PostgreSQL error: %s (Code: %s, Detail: %s)", pgErr.Message, pgErr.Code, pgErr.Detail),
			}
		}
		return "", fmt.Errorf("unexpected error: %w", err)
	}

	return user.String, nil
}
