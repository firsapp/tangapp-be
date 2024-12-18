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

func NewEventController(eventService *service.EventService) *EventController {
	return &EventController{eventService: eventService}
}

type EventMemberDetailRequest struct {
	EventID string `json:"event_id" binding:"required"`
	UserID  string `json:"user_id" binding:"required"`
}

type EventPurchaseDetailRequest struct {
	EventID    string `json:"event_id" binding:"required"`
	Name       string `json:"name" binding:"required"`
	Qty        int32  `json:"qty" binding:"required"`
	EachPrice  int32  `json:"each_price" binding:"required"`
	TotalPrice int32  `json:"total_price" binding:"required"`
}

type CreateEventRequest struct {
	CreatedBy            string                       `json:"created_by" binding:"required"`
	Title                string                       `json:"title" binding:"required"`
	Description          string                       `json:"description" binding:"max=60"`
	Status               string                       `json:"status"`
	TotalAmount          int32                        `json:"total_amount" binding:"required"`
	DateEvent            time.Time                    `json:"date_event" binding:"required"`
	CanEdit              bool                         `json:"can_edit"`
	EventMemberDetails   []EventMemberDetailRequest   `json:"event_member_details,omitempty"`   // Optional
	EventPurchaseDetails []EventPurchaseDetailRequest `json:"event_purchase_details,omitempty"` // Optional
}

func (ec *EventController) CreateEvent(ctx *gin.Context) {
	var req CreateEventRequest

	// Bind the incoming JSON to the struct (validation)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"validation error": err.Error()})
		return
	}

	// Map incoming request to EventPayload
	eventPayload := repository.EventPayload{
		CreatedBy:   req.CreatedBy,
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
		TotalAmount: req.TotalAmount,
		DateEvent:   req.DateEvent,
		CanEdit:     req.CanEdit,
	}

	// Insert []EventMemberDetailRequest to []EventMemberDetailPayload
	var memberDetails []repository.EventMemberDetailPayload
	for _, detail := range req.EventMemberDetails {
		memberDetails = append(memberDetails, repository.EventMemberDetailPayload{
			EventID: detail.EventID,
			UserID:  detail.UserID,
		})
	}

	// Insert []EventPurchaseDetailRequest to []EventPurchaseDetailPayload
	var purchaseDetails []repository.EventPurchaseDetailPayload
	for _, purchase := range req.EventPurchaseDetails {
		purchaseDetails = append(purchaseDetails, repository.EventPurchaseDetailPayload{
			EventID:    purchase.EventID,
			Name:       purchase.Name,
			Qty:        purchase.Qty,
			EachPrice:  purchase.EachPrice,
			TotalPrice: purchase.TotalPrice,
		})
	}

	// Call the service method to create the event
	err := ec.eventService.AddEvent(ctx, &eventPayload, &memberDetails, &purchaseDetails)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error while creating event": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "event created successfully"})
}
