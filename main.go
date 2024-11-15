package main

import (
	"tangapp-be/config"
	"tangapp-be/routes"

	"github.com/gin-gonic/gin"
)

// Intiates http server
func main() {
	config.LoadConfig() // Load configs

	r := gin.Default() // Gin router

	routes.SetupRoutes(r)

	r.Run(":8080")
}
