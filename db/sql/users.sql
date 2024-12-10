-- name: AddUser :one
INSERT INTO users (
  username,
  email
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: UpdateUser :one
UPDATE users
SET username = $2
WHERE id = $1
RETURNING *;
