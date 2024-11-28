// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: payment_histories.sql

package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const addPaymentHistory = `-- name: AddPaymentHistory :one
INSERT INTO payment_history
( event_member_details_id, from_user_id, to_user_id, nominal, description)
VALUES( $1, $2, $3, $4, $5)
RETURNING id, event_member_details_id, from_user_id, to_user_id, nominal, description, created_at
`

type AddPaymentHistoryParams struct {
	EventMemberDetailsID uuid.UUID      `json:"event_member_details_id"`
	FromUserID           uuid.UUID      `json:"from_user_id"`
	ToUserID             uuid.UUID      `json:"to_user_id"`
	Nominal              int32          `json:"nominal"`
	Description          sql.NullString `json:"description"`
}

func (q *Queries) AddPaymentHistory(ctx context.Context, arg AddPaymentHistoryParams) (PaymentHistory, error) {
	row := q.db.QueryRowContext(ctx, addPaymentHistory,
		arg.EventMemberDetailsID,
		arg.FromUserID,
		arg.ToUserID,
		arg.Nominal,
		arg.Description,
	)
	var i PaymentHistory
	err := row.Scan(
		&i.ID,
		&i.EventMemberDetailsID,
		&i.FromUserID,
		&i.ToUserID,
		&i.Nominal,
		&i.Description,
		&i.CreatedAt,
	)
	return i, err
}

const listPaymentHistoryByEvent = `-- name: ListPaymentHistoryByEvent :many
SELECT 
  ph.from_user_id,
  ph.to_user_id,
  ph.nominal, 
  ph.description, 
  ph.created_at 
FROM payment_history ph 
JOIN event_member_details emd on  ph.event_member_details_id = emd.id 
JOIN events e on emd.event_id = e.id  
WHERE e.id=$1
`

type ListPaymentHistoryByEventRow struct {
	FromUserID  uuid.UUID      `json:"from_user_id"`
	ToUserID    uuid.UUID      `json:"to_user_id"`
	Nominal     int32          `json:"nominal"`
	Description sql.NullString `json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
}

func (q *Queries) ListPaymentHistoryByEvent(ctx context.Context, eventID uuid.UUID) ([]ListPaymentHistoryByEventRow, error) {
	rows, err := q.db.QueryContext(ctx, listPaymentHistoryByEvent, eventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListPaymentHistoryByEventRow{}
	for rows.Next() {
		var i ListPaymentHistoryByEventRow
		if err := rows.Scan(
			&i.FromUserID,
			&i.ToUserID,
			&i.Nominal,
			&i.Description,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listPaymentHistoryByUser = `-- name: ListPaymentHistoryByUser :many
SELECT 
  e.id,
  ph.from_user_id,
  ph.to_user_id,
  ph.nominal, 
  ph.description, 
  ph.created_at 
FROM payment_history ph 
JOIN event_member_details emd on  ph.event_member_details_id = emd.id 
JOIN events e on emd.event_id = e.id  
WHERE ph.from_user_id=$1 OR ph.to_user_id=$2
`

type ListPaymentHistoryByUserParams struct {
	FromUserID uuid.UUID `json:"from_user_id"`
	ToUserID   uuid.UUID `json:"to_user_id"`
}

type ListPaymentHistoryByUserRow struct {
	ID          uuid.UUID      `json:"id"`
	FromUserID  uuid.UUID      `json:"from_user_id"`
	ToUserID    uuid.UUID      `json:"to_user_id"`
	Nominal     int32          `json:"nominal"`
	Description sql.NullString `json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
}

func (q *Queries) ListPaymentHistoryByUser(ctx context.Context, arg ListPaymentHistoryByUserParams) ([]ListPaymentHistoryByUserRow, error) {
	rows, err := q.db.QueryContext(ctx, listPaymentHistoryByUser, arg.FromUserID, arg.ToUserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListPaymentHistoryByUserRow{}
	for rows.Next() {
		var i ListPaymentHistoryByUserRow
		if err := rows.Scan(
			&i.ID,
			&i.FromUserID,
			&i.ToUserID,
			&i.Nominal,
			&i.Description,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updatePaymentHistory = `-- name: UpdatePaymentHistory :one
UPDATE payment_history
SET 
  event_member_details_id= $2,  
  to_user_id=$3, 
  nominal=$4,
  description= $5
WHERE id=$1
RETURNING id, event_member_details_id, from_user_id, to_user_id, nominal, description, created_at
`

type UpdatePaymentHistoryParams struct {
	ID                   uuid.UUID      `json:"id"`
	EventMemberDetailsID uuid.UUID      `json:"event_member_details_id"`
	ToUserID             uuid.UUID      `json:"to_user_id"`
	Nominal              int32          `json:"nominal"`
	Description          sql.NullString `json:"description"`
}

func (q *Queries) UpdatePaymentHistory(ctx context.Context, arg UpdatePaymentHistoryParams) (PaymentHistory, error) {
	row := q.db.QueryRowContext(ctx, updatePaymentHistory,
		arg.ID,
		arg.EventMemberDetailsID,
		arg.ToUserID,
		arg.Nominal,
		arg.Description,
	)
	var i PaymentHistory
	err := row.Scan(
		&i.ID,
		&i.EventMemberDetailsID,
		&i.FromUserID,
		&i.ToUserID,
		&i.Nominal,
		&i.Description,
		&i.CreatedAt,
	)
	return i, err
}
