package controllers

import (
	"fmt"
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

	c.Redirect(http.StatusFound, fmt.Sprintf("http://localhost:5173/loginsuccess?token=%s", token))

}
