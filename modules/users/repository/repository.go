package repository

import (
	"context"
	"database/sql"
	"tangapp-be/repository"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type UserRepository struct {
	q *repository.Queries
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		q: repository.New(db),
	}
}

// Adds a new user to the database
func (r *UserRepository) AddUser(ctx context.Context, username, email string) (repository.User, error) {
	arg := repository.AddUserParams{
		Username: sql.NullString{String: username, Valid: true},
		Email:    email,
	}

	user, err := r.q.AddUser(ctx, arg)
	if err != nil {
		return repository.User{}, err
	}
	return user, nil
}

// Gets a user by id (gak tau uuid.UUID bener atau enggak)
func (r *UserRepository) GetUserByID(ctx context.Context, id uuid.UUID) (repository.User, error) {
	return r.q.GetUser(ctx, id)
}
