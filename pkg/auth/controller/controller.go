package controller

import (
	"log"
	"net/http"
	"tangapp-be/config"
	"tangapp-be/queries"
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

func (ac *AuthController) GoogleAuthCallback(c *gin.Context) {
	oauth, err := gothic.CompleteUserAuth(c.Writer, c.Request) // Completes user auth
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	// Check if user exists
	user, exists, err := ac.authSvc.ValidateUserByEmail(c, oauth.Email)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to validate user"})
		return
	}

	// If user exists, Generate JWT
	var token string
	if exists {
		token, err = utils.GenerateJWT(user.ID.String(), user.Email, user.Username.String, config.JWTSecret)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate token"})
			return
		}
	} else {
		// If user does not exist, add user to database
		newUser, err := ac.authSvc.AddNewUser(c,
			queries.AddUserParams{
				Username: utils.ToNullString(oauth.Name),
				Email:    oauth.Email,
			})
		if err != nil {
			log.Print(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
			return
		}

		// Then, create the JWT
		token, err = utils.GenerateJWT(newUser.ID.String(), newUser.Email, newUser.Username.String, config.JWTSecret)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate token"})
			return
		}
	}

	// Redirect to frontend with token as a query parameter
	frontendCallbackURL := "http://localhost:5173/auth-callback" // Adjust as needed
	c.Redirect(http.StatusFound, frontendCallbackURL+"?token="+token)
}
