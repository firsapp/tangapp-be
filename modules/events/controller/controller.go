package controller

import (
	"net/http"
	"tangapp-be/modules/events/repository"
	"tangapp-be/modules/events/service"
	"time"

	"github.com/gin-gonic/gin"
)

type EventController struct {
	eventService *service.EventService
}

func NewUserController(eventService *service.EventService) *EventController {
	return &EventController{eventService: eventService}
}

var req struct {
	CreatedBy   string    `json:"created_by" binding:"required"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description" binding:"max=60"`
	Status      string    `json:"status"`
	TotalAmount int       `json:"total_amount"`
	DateEvent   time.Time `json:"date_event"`
	CanEdit     bool      `json:"can_edit"`
}

func (ec *EventController) CreateEvent(ctx *gin.Context) {

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"validation error": err.Error()})
		return
	}

	err := ec.eventService.CreateEvent(ctx, &repository.EventPayload{
		CreatedBy:   req.CreatedBy,
		Title:       req.Title,
		Description: req.Status,
		Status:      req.Status,
		TotalAmount: req.TotalAmount,
		DateEvent:   req.DateEvent,
		CanEdit:     req.CanEdit,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error while creating event :": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "event created successfully"})
}
