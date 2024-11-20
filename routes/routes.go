package routes

import (
	"net/http"
	"tangapp-be/controllers"
	"tangapp-be/middleware"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
)

var baseUri = "/api/v1"

func SetupRoutes(r *gin.Engine) {
	// Google Oauth routes
	r.GET(baseUri+"/auth/:provider", func(c *gin.Context) {
		gothic.GetProviderName = func(req *http.Request) (string, error) {
			return "google", nil // specify the provider (e.g., Google)
		}
		gothic.BeginAuthHandler(c.Writer, c.Request)
	})
	r.GET("/auth/:provider/callback", controllers.GoogleAuthCallback)

	// Protected routes
	protected := r.Group("/protected")
	protected.Use(middleware.JWTAuthMiddleware())
	protected.GET("/profile", func(c *gin.Context) {
		userID, _ := c.Get("userID")
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to your profile!", "user_id": userID})
	})

}
