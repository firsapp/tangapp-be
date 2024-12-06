package service

import (
	"context"
	"tangapp-be/modules/users/repository"
	query "tangapp-be/repository"
)

type UserService struct {
	r *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{r: repo}
}

// Handles user creation logic
func (s *UserService) CreateUser(ctx context.Context, username, email string) (query.User, error) {

	// Business lojig

	user, err := s.r.AddUser(ctx, username, email)
	if err != nil {
		return query.User{}, err
	}
	return user, nil
}
