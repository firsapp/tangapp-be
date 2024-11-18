package controllers

import (
	"net/http"
	"tangapp-be/config"
	"tangapp-be/utils"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
)

func GoogleAuthHandler(c *gin.Context) {
	gothic.BeginAuthHandler(c.Writer, c.Request) // Starts authentication process
}

func GoogleAuthCallback(c *gin.Context) {
	user, err := gothic.CompleteUserAuth(c.Writer, c.Request) // Completes user auth
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

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
