package controllers

import (
	"net/http"
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
	}

	//  Generate JWT
	token, err := uti
}
