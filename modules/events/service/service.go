package service

import (
	"context"
	"tangapp-be/modules/events/repository"
)

type EventService struct {
	r *repository.EventRepository
}

func NewEventService(repo *repository.EventRepository) *EventService {
	return &EventService{r: repo}
}

func (s *EventService) AddEvent(ctx context.Context, arg *repository.EventPayload, memberDetails *[]repository.EventMemberDetailPayload, purchaseDetails *[]repository.EventPurchaseDetailPayload) error {
	// Start a transaction
	tx, err := s.r.BeginTransaction(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback(ctx) // Rollback the transaction if an error occurs
		}
	}()

	// Create the event using transaction-bound queries
	err = s.r.AddEvent(ctx, arg)
	if err != nil {
		return err
	}

	// If member details are provided, insert them
	if len(*memberDetails) > 0 {
		for _, member := range *memberDetails {
			err = s.r.AddEventMemberDetail(ctx, &member)
			if err != nil {
				return err
			}
		}
	}

	// If purchase details are provided, insert them
	if len(*purchaseDetails) > 0 {
		for _, purchase := range *purchaseDetails {
			err = s.r.AddEventPurchaseDetail(ctx, &purchase)
			if err != nil {
				return err
			}
		}
	}

	// If everything is successful, commit the transaction
	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}
