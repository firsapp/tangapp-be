package repository

import (
	"context"
	"tangapp-be/queries"
	"tangapp-be/utils"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4/pgxpool"
)

type EventRepository struct {
	q *queries.Queries
}

var pgErr *pgconn.PgError

func NewEventRepository(db *pgxpool.Pool) *EventRepository {
	return &EventRepository{
		q: queries.New(db),
	}
}

type EventPayload struct {
	CreatedBy   string
	Title       string
	Description string
	Status      string
	TotalAmount int
	DateEvent   time.Time
	CanEdit     bool
}

func (r *EventRepository) CreateEvent(ctx context.Context, arg *EventPayload) error {
	_, err := r.q.AddEvent(ctx, queries.AddEventParams{
		CreatedBy:   uuid.MustParse(arg.CreatedBy),
		Title:       utils.ToNullString(arg.Title),
		Description: utils.ToNullString(arg.Description),
		Status:      arg.Status,
		TotalAmount: int32(arg.TotalAmount),
		DateEvent:   utils.ToNullTime(arg.DateEvent),
		CanEdit:     arg.CanEdit,
	})
	if err != nil {
		return err
	}

	return nil
}
