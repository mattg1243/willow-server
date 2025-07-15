-- name: CreatePaymentType :one
INSERT INTO payment_types (
  user_id, "name"
) VALUES (
  $1, $2
) RETURNING *;

-- name: GetPaymentTypes :many
SELECT * 
FROM payment_types
WHERE user_id = $1 OR NULL;

-- name: UpdatePaymentType :one
UPDATE payment_types
SET
  "name" = $2
WHERE
  id = $1
RETURNING *;

-- name: DeletePaymentType :exec
DELETE FROM payment_types WHERE id = $1;