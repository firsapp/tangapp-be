package config

import (
	"log"

	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
)

func LoadOauthConfig(config Config) {
	callbackUri := "http://localhost:8080/auth/google/callback"

	config, err := LoadConfig("../..") // ../.. means "go to parent folder"
	if err != nil {
		log.Fatal("oauth : can't load configuration files")
	}

	goth.UseProviders(
		google.New(
			config.GoogleClientID,
			config.GoogleClientSecret,
			callbackUri, "email", "profile",
		),
	)
}
