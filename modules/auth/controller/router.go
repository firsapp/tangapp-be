package controller

import (
	"tangapp-be/pkg/auth/service"

	"github.com/gin-gonic/gin"
)

// AuthController defines a handler with the required dependencies.
type AuthController struct {
	authSvc service.AuthService
}

// NewAuthController returns an instance of TodoHandler.
func NewAuthController(authSvc service.AuthService) *AuthController {
	return &AuthController{
		authSvc: authSvc,
	}
}

// Register registers the HTTP REST handlers route.
func (h *AuthController) Register(r *gin.Engine) {

	auth := r.Group("/v1/auth")

	auth.GET("/:provider", h.GoogleAuthHandler)
	auth.GET("/:provider/callback", h.GoogleAuthCallback)

}
