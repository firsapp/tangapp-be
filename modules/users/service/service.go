package service

import (
	"context"
	"log"
	"tangapp-be/errors"
	"tangapp-be/modules/users/repository"
	"tangapp-be/queries"

	"github.com/google/uuid"
	"github.com/jackc/pgx"
)

type UserService struct {
	r *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{r: repo}
}

// Handles user creation logic
func (s *UserService) CreateUser(ctx context.Context, username, email string) (queries.User, error) {

	// Business lojig

	user, err := s.r.AddUser(ctx, username, email)
	if err != nil {
		return queries.User{}, err
	}

	return user, nil
}

func (s *UserService) GetUserByID(ctx context.Context, ID uuid.UUID) (queries.User, error) {

	user, err := s.r.GetUserByID(ctx, ID)
	if err != nil {
		log.Printf("Error type: %T", err) // Add this to check the error type
		if err == pgx.ErrNoRows {
			// User not found
			return queries.User{}, &errors.UserNotFoundError{ID: ID}
		}
		return queries.User{}, err
	}
	return user, nil
}
