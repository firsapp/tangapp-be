// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: user.sql

package repository

import (
	"context"
	"database/sql"
)

const createAccount = `-- name: CreateAccount :one
INSERT INTO users (
  name,
  title,
  created_at
) VALUES (
  $1, $2, $3
)
RETURNING id, name, title, created_at
`

type CreateAccountParams struct {
	Name      sql.NullString `json:"name"`
	Title     sql.NullString `json:"title"`
	CreatedAt sql.NullTime   `json:"created_at"`
}

func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createAccount, arg.Name, arg.Title, arg.CreatedAt)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Title,
		&i.CreatedAt,
	)
	return i, err
}

const getAccount = `-- name: GetAccount :one
SELECT id, name, title, created_at FROM users
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetAccount(ctx context.Context, id int32) (User, error) {
	row := q.db.QueryRowContext(ctx, getAccount, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Title,
		&i.CreatedAt,
	)
	return i, err
}
