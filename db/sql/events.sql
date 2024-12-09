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
WHERE id = $1 AND is_active = true LIMIT 1;

-- name: ListEventByUser :many
SELECT * FROM events
WHERE created_by = $1 AND is_active = true;

-- name: UpdateEvent :one
UPDATE events
SET 
  title = $2,
  description = $3,
  status = $4,
  total_amount = $5,
  date_event = $6
WHERE id = $1 AND can_edit = true AND is_active = true
RETURNING *;

-- name: DeleteEvent :one
UPDATE events
SET 
  is_active = $2
WHERE id = $1 AND can_edit = true
RETURNING *;