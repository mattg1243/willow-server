-- name: CreatePaymentType :one
INSERT INTO payment_types (
  user_id, "name"
) VALUES (
  $1, $2
) RETURNING *;

-- name: GetPaymentTypes :many
SELECT * 
FROM payment_types
WHERE user_id = $1 OR $1 IS NULL;

-- name: GetDefaultPaymentTypes :many
SELECT *
FROM payment_types
WHERE user_id IS NULL;

-- name: GetPaymentType :one
SELECT *
FROM payment_types
WHERE id = $1 AND user_id = $2;

-- name: UpdatePaymentType :one
UPDATE payment_types
SET
  "name" = $3
WHERE
  user_id = $1 AND id = $2
RETURNING *;

-- name: DeletePaymentType :exec
DELETE FROM payment_types WHERE user_id = $1 AND id = $2;