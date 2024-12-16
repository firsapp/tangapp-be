package router

import (
	"tangapp-be/modules/events/controller"

	"github.com/gin-gonic/gin"
)

func RegisterEventRoutes(r *gin.Engine, eventController *controller.EventController) {
	events := r.Group("/v1/events")
	{
		events.POST("/", eventController.CreateEvent)

	}
}
