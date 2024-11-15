package main

import (
	"tangapp-be/config"

	"github.com/gin-gonic/gin"
)

// Intiates http server
func main() {
	config.LoadConfig() // Load configs

	r := gin.Default() // Gin router

}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
