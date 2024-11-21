-- name: AddEvent :one
INSERT INTO events (
  created_by,
  title,
  description,
  status,
  total_amount,
  date_event,
  can_edit
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
)
RETURNING *;

-- name: GetEvent :one
SELECT * FROM events
WHERE id = $1 LIMIT 1;

-- name: GetEventByUser :many
SELECT * FROM events
WHERE created_by = $1;

-- name: UpdateEvent :one
-- UPDATE events
-- SET username = $2
-- WHERE id = $1
-- RETURNING *;