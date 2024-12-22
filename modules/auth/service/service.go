package service

import (
	"context"
	"fmt"
	"log"
	"tangapp-be/config"
	"tangapp-be/modules/auth/repository"
	"tangapp-be/utils"

	"github.com/markbates/goth"
)

type AuthService interface {
	GoogleAuthCallbackHandler(ctx context.Context, arg goth.User) (string, error)
}

type authService struct {
	repo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) AuthService {
	return &authService{repo: repo}
}

func (s *authService) GoogleAuthCallbackHandler(ctx context.Context, arg goth.User) (string, error) {

	// Check if user exists
	user, exists, err := s.repo.ValidateUserByEmail(ctx, arg.Email)
	if err != nil {
		log.Print(err)
		return "", fmt.Errorf("failed to validate user : %w", err)
	}

	// If user exists, Generate JWT
	var token string
	if exists {
		token, err = utils.GenerateJWT(user.ID.String(), user.Email, user.Username.String, config.JWTSecret)
		if err != nil {
			return "", fmt.Errorf("failed to generate token : %w", err)
		}
	} else {
		// If user does not exist, add user to database
		newUser, err := s.repo.AddNewUser(ctx,
			repository.AddNewUserPayload{
				Username: arg.Name,
				Email:    arg.Email,
			})
		if err != nil {
			log.Print(err)
			return "", fmt.Errorf("failed to create user : %w", err)
		}

		// Then, create the JWT
		token, err = utils.GenerateJWT(newUser.ID.String(), newUser.Email, newUser.Username.String, config.JWTSecret)
		if err != nil {
			return "", fmt.Errorf("failed to genreate token : %w", err)
		}
	}

	return token, nil
}
