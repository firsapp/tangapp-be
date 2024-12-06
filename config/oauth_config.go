package config

import (
	"fmt"

	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
)

// LoadOauthConfig used for generate google oauth config
func LoadOauthConfig(config Config) {
	callbackUrl := fmt.Sprintf("http://%s/v1/auth/google/callback", BaseUrl)

	config, err := LoadConfig("../") // ../.. means "go to parent folder"
	if err != nil {
		println("unable to load configuration file")
	}

	goth.UseProviders(
		google.New(
			config.GoogleClientID,
			config.GoogleClientSecret,
			callbackUrl, "email", "profile",
		),
	)

	println("Google OAuth provider initialized with callback:", callbackUrl)
}
