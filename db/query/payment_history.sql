-- name: AddPaymentHistory :one
INSERT INTO payment_history
( event_member_details_id, to_user_id, nominal, description, created_at)
VALUES( $1, $2, $3, $4, $5)
RETURNING *;

-- name: ListPaymentHistory :many
SELECT 
  ph.to_user_id,
  ph.nominal, 
  ph.description, 
  ph.created_at 
FROM payment_history ph 
JOIN event_member_details emd on  ph.event_member_details_id = emd.id 
JOIN events e on emd.event_id = e.id  
WHERE e.id=$1;

-- name: UpdatePaymentHistory :one
UPDATE payment_history
SET 
  event_member_details_id= $2, 
  to_user_id=$3, 
  nominal=$4,
  description= $5, 
  created_at= $6
WHERE id=$1
RETURNING *;