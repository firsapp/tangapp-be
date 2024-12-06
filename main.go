package main

import (
	"log"
	"tangapp-be/config"

	authController "tangapp-be/pkg/auth/controller"
	authRepository "tangapp-be/pkg/auth/repository"
	authService "tangapp-be/pkg/auth/service"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Config struct {
	DB *pgxpool.Pool
}

// Intiates http server
func main() {
	cfg := new(Config)
	configuration, err := config.LoadConfig(".") // Load configs
	if err != nil {
		log.Fatal(err)
	}
	config.LoadOauthConfig(configuration)
	r := gin.Default() // Gin router

	// routes.SetupRoutes(r)

	authRepo := authRepository.NewAuth(cfg.DB)
	authSvc := authService.NewAuth(authRepo)
	authController.NewAuthController(authSvc).Register(r)

	r.Run(config.BaseUrl)

}
