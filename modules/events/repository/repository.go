package repository

import (
	"context"
	"tangapp-be/queries"
	"tangapp-be/utils"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type EventRepository struct {
	q  *queries.Queries
	db *pgxpool.Pool
}

func NewEventRepository(db *pgxpool.Pool) *EventRepository {
	return &EventRepository{
		q:  queries.New(db),
		db: db,
	}
}

type EventPayload struct {
	ID          string
	Title       string
	Description string
	Status      string
	TotalAmount int32
	DateEvent   time.Time
	CreatedBy   string
	CanEdit     bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	IsActive    bool
}

type EventMemberDetailPayload struct {
	ID           string
	EventID      string
	UserID       string
	Bill         int32
	Paid         int32
	Compensation int32
	Notes        string
	Done         bool
}

type EventPurchaseDetailPayload struct {
	ID         string
	EventID    string
	Name       string
	Qty        int32
	EachPrice  int32
	TotalPrice int32
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// Starts a transaction
func (r *EventRepository) BeginTransaction(ctx context.Context) (*pgxpool.Tx, *queries.Queries, error) {
	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, nil, err
	}
	qtx := r.q.WithTx(tx)

	return tx.(*pgxpool.Tx), qtx, nil
}

func (r *EventRepository) AddEvent(ctx context.Context, qtx *queries.Queries, arg *EventPayload) error {
	_, err := qtx.AddEvent(ctx, queries.AddEventParams{
		CreatedBy:   uuid.MustParse(arg.CreatedBy),
		Title:       utils.ToNullString(arg.Title),
		Description: utils.ToNullString(arg.Description),
		Status:      arg.Status,
		TotalAmount: arg.TotalAmount,
		DateEvent:   utils.ToNullTime(arg.DateEvent),
		CanEdit:     arg.CanEdit,
	})
	if err != nil {
		return err
	}

	return nil
}

func (r *EventRepository) AddEventPurchaseDetail(ctx context.Context, qtx *queries.Queries, arg *EventPurchaseDetailPayload) error {
	_, err := qtx.AddPurchaseDetail(ctx, queries.AddPurchaseDetailParams{
		EventID:    uuid.MustParse(arg.EventID),
		Name:       arg.Name,
		Qty:        arg.Qty,
		EachPrice:  arg.EachPrice,
		TotalPrice: arg.TotalPrice,
	})
	if err != nil {
		return err
	}

	return nil
}

func (r *EventRepository) AddEventMemberDetail(ctx context.Context, qtx *queries.Queries, arg *EventMemberDetailPayload) error {
	_, err := qtx.AddMemberDetail(ctx, queries.AddMemberDetailParams{
		EventID:      uuid.MustParse(arg.EventID),
		UserID:       uuid.MustParse(arg.UserID),
		Bill:         utils.ToNullInt32(arg.Bill),
		Paid:         utils.ToNullInt32(arg.Paid),
		Compensation: utils.ToNullInt32(arg.Compensation),
		Notes:        utils.ToNullString(arg.Notes),
		Done:         arg.Done,
	})
	if err != nil {
		return err
	}

	return nil
}
