package config

import (
	"log"

	"github.com/spf13/viper"
)

var JWTSecret []byte

type Config struct {
	JWTSecret          string `mapstructure:"JWT_SECRET"`
	DBCredential       string `mapstructure:"DB_CREDENTIAL"`
	GoogleClientID     string `mapstructure:"GOOGLE_CLIENT_ID"`
	GoogleClientSecret string `mapstructure:"GOOGLE_CLIENT_SECRET"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal("fatal: unable to load configurations")
	}

	err = viper.Unmarshal(&config)
	return
}
