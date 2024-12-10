package service

import (
	"context"
	"tangapp-be/pkg/auth/repository"
	"tangapp-be/queries"
)

type AuthService interface {
	ValidateUserByEmail(ctx context.Context, email string) (queries.User, bool, error)
	AddNewUser(ctx context.Context, user queries.AddUserParams) (queries.User, error)
}

type authService struct {
	repo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) AuthService {
	return &authService{repo: repo}
}

func (s *authService) ValidateUserByEmail(ctx context.Context, email string) (queries.User, bool, error) {
	user, exists, err := s.repo.ValidateUserByEmail(ctx, email)
	if err != nil {
		return queries.User{}, exists, err // This is the case for unknown error
	}
	return user, exists, nil // This belongs to both if user exists and not exists case
}

func (s *authService) AddNewUser(ctx context.Context, arg queries.AddUserParams) (queries.User, error) {
	user, err := s.repo.AddNewUser(ctx, arg)
	if err != nil {
		return queries.User{}, err
	}
	return user, nil
}
