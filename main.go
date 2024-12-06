package main

import (
	"context"
	"log"
	"tangapp-be/config"
	"tangapp-be/modules/users/controller"
	"tangapp-be/modules/users/repository"
	"tangapp-be/modules/users/router"
	"tangapp-be/modules/users/service"

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

	//Experimental - database
	connString := configuration.DBCredential
	db, err := pgxpool.Connect(context.Background(), connString)
	if err != nil {
		log.Fatal("kenot konek db")
	}

	// auth
	authRepo := authRepository.NewAuth(cfg.DB)
	authSvc := authService.NewAuth(authRepo)
	authController.NewAuthController(authSvc).Register(r)

	// users
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)
	router.RegisterUserRoutes(r, userController)

	r.Run(config.BaseUrl)

}
