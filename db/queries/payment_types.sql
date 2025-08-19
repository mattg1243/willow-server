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

-- name: GetPaymentTypeByEvent :one
SELECT pt.name
FROM events_payment_types ept
JOIN payment_types pt ON pt.id = ept.payment_type_id
WHERE ept.event_id = $1 AND pt.user_id = $2;

-- name: UpdatePaymentType :one
UPDATE payment_types
SET
  "name" = $3
WHERE
  user_id = $1 AND id = $2
RETURNING *;

-- name: DeletePaymentType :exec
DELETE FROM payment_types WHERE user_id = $1 AND id = $2;

-- name: AddPaymentTypeToEvent :exec
INSERT INTO events_payment_types (
  event_id, payment_type_id
) VALUES (
  $1, $2
);

-- name: RemovePaymentTypeFromEvent :exec
DELETE FROM events_payment_types WHERE event_id = $1;
