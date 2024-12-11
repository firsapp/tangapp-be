package router

import (
	"tangapp-be/modules/users/controller"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine, userController *controller.UserController) {
	users := r.Group("/v1/users")
	{
		users.POST("/", userController.CreateUser)
		users.GET("/", userController.GetUserByID)
		users.PATCH("/", userController.UpdateUser)
	}
}
