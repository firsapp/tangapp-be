package config

import (
	"fmt"
	"log"

	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
)

// LoadOauthConfig used for generate google oauth config
func LoadOauthConfig(config Config) {
	callbackUri := fmt.Sprintf("%s/auth/google/callback", BaseUrl)

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
