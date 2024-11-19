package main

import (
	"log"
	"tangapp-be/config"
	"tangapp-be/routes"

	"github.com/gin-gonic/gin"
)

// Intiates http server
func main() {
	configuration, err := config.LoadConfig(".") // Load configs
	if err != nil {
		log.Fatal(err)
	}
	config.LoadOauthConfig(configuration)
	r := gin.Default() // Gin router

	routes.SetupRoutes(r)

	r.Run("127.0.0.1:8080")
}
