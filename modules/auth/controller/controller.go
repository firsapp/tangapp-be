package controller

import (
	"net/http"
	"tangapp-be/config"
	"tangapp-be/utils"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
)

func (h *AuthController) GoogleAuthHandler(c *gin.Context) {
	gothic.GetProviderName = func(req *http.Request) (string, error) {
		return c.Param("provider"), nil // Specify the provider (e.g., Google)
	}
	gothic.BeginAuthHandler(c.Writer, c.Request) // Starts authentication process
}

func (h *AuthController) GoogleAuthCallback(c *gin.Context) {
	user, err := gothic.CompleteUserAuth(c.Writer, c.Request) // Completes user auth
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	// Check if user exists

	// Generate JWT
	token, err := utils.GenerateJWT(user.UserID, user.Email, user.Name, config.JWTSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate token"})
		return
	}

	// Redirect to frontend with token as a query parameter
	frontendCallbackURL := "http://localhost:5173/auth-callback" // Adjust as needed
	c.Redirect(http.StatusFound, frontendCallbackURL+"?token="+token)
}