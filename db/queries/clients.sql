-- name: CreateClient :one
INSERT INTO clients (
  user_id, fname, lname, email, phone, rate, balanceNotifyThreshold, id, created_at, updated_at
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, NOW(), NOW()
) RETURNING *;

-- name: GetClient :one
SELECT *
FROM clients
WHERE id = $1;

-- name: GetClients :many
SELECT *
FROM clients
WHERE clients.user_id = $1;

-- name: UpdateClient :exec
UPDATE clients
SET
    fname = $2,
    lname = $3,
    email = $4,
    balance = $5,
    balancenotifythreshold = $6,
    rate = $7,
    isarchived = $8,
    update_at = NOW()
WHERE
    id = $1;

-- name: DeleteClient :exec
DELETE FROM clients
WHERE id = $1;
