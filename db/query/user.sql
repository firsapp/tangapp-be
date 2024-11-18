-- name: CreateAccount :one
INSERT INTO users (
  name,
  title,
  created_at
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetAccount :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;