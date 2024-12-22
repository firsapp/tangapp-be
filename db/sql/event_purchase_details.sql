-- name: AddPurchaseDetail :one
INSERT INTO event_purchase_details
( event_id, name, qty, each_price, total_price)
VALUES($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetPurchaseDetailByID :one
SELECT id, event_id, name, qty, each_price, total_price, updated_at
FROM event_purchase_details
WHERE id = $1;

-- name: GetPurchaseDetailByEventID :many
SELECT id, event_id, name, qty, each_price, total_price, updated_at
FROM event_purchase_details
WHERE event_id = $1;

-- name: UpdatePurchaseDetail :one
UPDATE event_purchase_details
SET
  name = $2,
  qty = $3,
  each_price = $4,
  total_price = $5,
  updated_at = $6
WHERE id = $1
RETURNING *;

-- name: DeletePurchaseDetail :exec