-- name: CreateEvent :one
INSERT INTO events (
    client_id, user_id, date, duration, event_type_id, detail, rate, amount, running_balance, id
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
) RETURNING *;

-- name: GetEvent :one
SELECT 
    e.id as id,
    e.client_id as client_id,
    e.user_id as user_id,
    e.date::timestamptz as "date",
    e.duration as duration,
    et.id as event_type_id,
    e.detail as detail,
    e.rate as rate,
    e.amount::INTEGER as amount,
    e.running_balance::INTEGER as running_balance,
    et.charge as charge
FROM events e
INNER JOIN event_types et ON e.event_type_id = et.id
WHERE e.id = $1;

-- name: GetEvents :many
SELECT 
    e.id as id,
    e.user_id as user_id,
    e.client_id as client_id,
    e.date::timestamptz as "date",
    e.duration as duration,
    et.id as event_type_id,
    e.detail as detail,
    e.rate as rate,
    e.amount::INTEGER as amount,
    e.running_balance::INTEGER as running_balance,
    et.charge as charge
FROM events e
INNER JOIN event_types et ON e.event_type_id = et.id
WHERE e.client_id = $1 or e.user_id = $1
ORDER BY e.date ASC;

-- name: UpdateEvent :one
UPDATE events
SET
    date = $2,
    duration = $3,
    event_type_id = $4,
    detail = $5,
    rate = $6,
    amount = $7,
    running_balance = $8
WHERE
    id = $1
RETURNING *;

-- name: UpdateRunningBalance :exec
UPDATE events
SET
    running_balance = $2
WHERE
    id = $1;

-- name: DeleteEvent :exec
DELETE FROM events
WHERE id = $1;