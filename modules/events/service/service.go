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

func (s *EventService) CreateEvent(ctx context.Context, arg *repository.EventPayload) error {
	err := s.r.CreateEvent(ctx, arg)
	if err != nil {
		return err
	}

	return nil

}
