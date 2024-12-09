package controller

import (
	"log"
	"net/http"
	"tangapp-be/errors"
	"tangapp-be/modules/users/service"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{userService: userService}
}

type UserResponse struct {
	ID        uuid.UUID
	Username  string
	Email     string
	CreatedAt time.Time
}

func (uc *UserController) CreateUser(c *gin.Context) {
	// To-do : haruskah bikin struct dari masing2 request ? atau bikin local kayak gini aja
	var req struct {
		Username string `json:"username" binding:"required,alphanum"`
		Email    string `json:"email" binding:"required,email"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := uc.userService.CreateUser(c, req.Username, req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (uc *UserController) GetUserByID(c *gin.Context) {
	var req struct {
		ID uuid.UUID `json:"id" binding:"required,uuid"`
	}

	// validasi request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data", "details": err.Error()})
		return
	}

	user, err := uc.userService.GetUserByID(c, req.ID)
	if err != nil {
		log.Printf("Error type in controller: %T", err)

		switch e := err.(type) {
		case *errors.UserNotFoundError:
			// If user not found, return not found error
			c.JSON(http.StatusNotFound, gin.H{"error": e.Error()})
		default:
			// If the error is unknown, return server error
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error", "details": e.Error()})
		}
		return
	}

	res := UserResponse{
		ID:        user.ID,
		Username:  user.Username.String, //Assert directly to .String aja lah
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	c.JSON(http.StatusOK, res)
}
