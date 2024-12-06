package controller

import "github.com/gin-gonic/gin"

// AuthService defines a contract for the handlers to be implemented.
type AuthService interface {
}

// AuthController defines a handler with the required dependencies.
type AuthController struct {
	authSvc AuthService
}

// NewAuthController returns an instance of TodoHandler.
func NewAuthController(authSvc AuthService) *AuthController {
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
