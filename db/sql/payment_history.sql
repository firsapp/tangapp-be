-- name: AddPaymentHistory :one
INSERT INTO payment_history
( event_member_details_id, from_user_id, to_user_id, nominal, description)
VALUES( $1, $2, $3, $4, $5)
RETURNING *;

-- name: ListPaymentHistoryByEvent :many
SELECT 
  ph.from_user_id,
  ph.to_user_id,
  ph.nominal, 
  ph.description, 
  ph.created_at 
FROM payment_history ph 
JOIN event_member_details emd on  ph.event_member_details_id = emd.id 
JOIN events e on emd.event_id = e.id  
WHERE e.id=sqlc.arg(event_id);

-- name: ListPaymentHistoryByUser :many
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
WHERE ph.from_user_id=$1 OR ph.to_user_id=$2;


-- name: UpdatePaymentHistory :one
UPDATE payment_history
SET 
  event_member_details_id= $2,  
  to_user_id=$3, 
  nominal=$4,
  description= $5
WHERE id=$1
RETURNING *;