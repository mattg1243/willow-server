-- name: CreateEvent :one
INSERT INTO events (
    client_id, date, duration, event_type_id, detail, rate, amount, newbalance, id
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9
) RETURNING *;

-- name: GetEvent :one
SELECT *
FROM events
WHERE id = $1;

-- name: GetEvents :many
SELECT *
FROM events
WHERE events.client_id = $1;

-- name: UpdateEvent :exec
UPDATE events
SET
    date = $2,
    duration = $3,
    event_type_id = $4,
    detail = $5,
    rate = $6,
    amount = $7,
    newbalance = $8
WHERE
    id = $1;

-- name: DeleteEvent :exec
DELETE FROM events
WHERE id = $1;
