-- name: AddMemberDetail :one
INSERT INTO event_member_details
(event_id, user_id, bill, paid, compensation, notes, done)
VALUES(
  $1, $2, $3, $4, $5, $6 ,$7
)
RETURNING *;

-- name: GetMemberDetail :one
SELECT 
  id, 
  event_id, 
  user_id, 
  bill, 
  paid, 
  compensation, 
  notes, 
  done
FROM event_member_details
WHERE id=$1;

-- name: ListMemberDetail :many
SELECT 
  id, 
  event_id, 
  user_id, 
  bill, 
  paid, 
  compensation, 
  notes, 
  done
FROM event_member_details
WHERE event_id=$1;

-- name: UpdateMemberDetail :one
UPDATE event_member_details 
SET 
  bill = $2, 
  paid = $3, 
  compensation =$4, 
  notes=$5, 
  done=$6
WHERE id = $1
RETURNING *;
