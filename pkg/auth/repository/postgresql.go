package repository

import (
	"context"
	"fmt"
	"tangapp-be/queries"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

// Harusnya disini kocak, bukan di service wleeeek
type AuthRepository interface {
	ValidateUserByEmail(ctx context.Context, email string) (queries.User, bool, error)
	AddNewUser(ctx context.Context, user queries.AddUserParams) (queries.User, error)
}

type authRepository struct {
	db *pgxpool.Pool
	q  *queries.Queries
}

func NewAuthRepository(db *pgxpool.Pool) AuthRepository {
	return &authRepository{
		db: db,
		q:  queries.New(db),
	}
}

func (a *authRepository) ValidateUserByEmail(ctx context.Context, email string) (queries.User, bool, error) {
	user, err := a.q.GetUserByEmail(ctx, email)
	if err != nil {
		if err == pgx.ErrNoRows {
			return queries.User{}, false, nil // User does not exist
		}
		return queries.User{}, false, fmt.Errorf("database error while validating user by email: %w", err)
	}
	return user, true, nil
}

func (a *authRepository) AddNewUser(ctx context.Context, arg queries.AddUserParams) (queries.User, error) {
	user, err := a.q.AddUser(ctx, arg)
	if err != nil {
		return queries.User{}, fmt.Errorf("database error while adding new user: %w", err)
	}
	return user, nil
}
