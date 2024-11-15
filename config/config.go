package config

import (
	"log"
	"os"

	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
)

var JWTSecret []byte

func LoadConfig() {
	// Loads JWT secret from .env
	JWTSecret = []byte(os.Getenv("JWT_SECRET"))
	if len(JWTSecret) == 0 {
		log.Fatal("fatal : unable to load jwt secret")
	}

	callbackUri := "http://localhost:8080/auth/google/callback"
	// Initialize google oauth credentials
	goth.UseProviders(
		google.New(
			os.Getenv("GOOGLE_CLIENT_ID"),
			os.Getenv("GOOGLE_CLIENT_SECRET"),
			callbackUri, "email", "profile",
		),
	)
}
