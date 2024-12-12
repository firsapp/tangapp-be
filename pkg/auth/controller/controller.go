package controller

import (
	"net/http"
	"tangapp-be/utils"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
)

func (h *AuthController) GoogleAuthHandler(ctx *gin.Context) {
	gothic.GetProviderName = func(req *http.Request) (string, error) {
		return ctx.Param("provider"), nil // Specify the provider (e.g., Google)
	}
	gothic.BeginAuthHandler(ctx.Writer, ctx.Request) // Starts authentication process
}

func (ac *AuthController) GoogleAuthCallback(ctx *gin.Context) {
	oauth, err := gothic.CompleteUserAuth(ctx.Writer, ctx.Request) // Completes user auth
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	token, err := ac.authSvc.GoogleAuthCallbackHandler(ctx, oauth)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
	}

	// Redirect to frontend with token as a query parameter
	frontendCallbackURL := "http://localhost:5173/auth-callback" // Adjust as needed
	ctx.Redirect(http.StatusFound, frontendCallbackURL+"?token="+token)
}
