-- name: CreateEvent :one
INSERT INTO events (
    client_id, user_id, date, duration, event_type_id, detail, rate, amount, running_balance, id, created_at, updated_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, NOW(), NOW()
) RETURNING *;

-- name: GetEvent :one
SELECT 
    e.id as id,
    e.client_id as client_id,
    e.user_id as user_id,
    e.date::timestamptz as "date",
    e.duration as duration,
    et.id as event_type_id,
    et.title as event_type_title,
    e.detail as detail,
    e.rate as rate,
    e.amount::INTEGER as amount,
    e.running_balance::INTEGER as running_balance,
    e.paid as paid,
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
    et.title as event_type_title,
    e.detail as detail,
    e.rate as rate,
    e.amount::INTEGER as amount,
    e.running_balance::INTEGER as running_balance,
    e.paid as paid,
    et.charge as charge
FROM events e
INNER JOIN event_types et ON e.event_type_id = et.id
WHERE e.client_id = $1 or e.user_id = $1
    AND ($2::timestamptz IS NULL OR e.date >= $2::timestamptz)
    AND ($3::timestamptz IS NULL OR e.date <= $3::timestamptz)
ORDER BY e.date ASC;

-- name: SetEventsToMiscType :exec
UPDATE events
SET event_type_id = $2
WHERE event_type_id = $1;

-- name: GetPayoutEvents :many
SELECT 
    e.id as id,
    e.user_id as user_id,
    e.client_id as client_id,
    e.date::timestamptz as "date",
    e.duration as duration,
    et.id as event_type_id,
    et.title as event_type_title,
    e.detail as detail,
    e.rate as rate,
    e.amount::INTEGER as amount,
    e.running_balance::INTEGER as running_balance,
    e.paid as paid,
    et.charge as charge
FROM events e
INNER JOIN event_types et ON e.event_type_id = et.id
INNER JOIN payout_events pe ON pe.event_id = e.id
WHERE pe.payout_id = $1
ORDER BY e.date ASC;

-- name: EventIsInPayout :one
SELECT EXISTS (
    SELECT 1
    FROM payout_events pe
    JOIN payouts p on pe.payout_id = p.id
    WHERE pe.event_id = ANY($1::uuid[]) AND p.user_id = $2
) AS is_in_user_payouts;

-- name: UpdateEvent :one
UPDATE events
SET
    date = $2,
    duration = $3,
    event_type_id = $4,
    detail = $5,
    rate = $6,
    amount = $7,
    running_balance = $8,
    paid = $9,
    updated_at = NOW()
WHERE
    id = $1
RETURNING *;

-- name: MarkEventPaid :one
UPDATE events
SET 
    paid = $2,
    updated_at = NOW()

WHERE
    id = $1
RETURNING *;

-- name: UpdateRunningBalance :exec
UPDATE events
SET
    running_balance = $2
WHERE
    id = $1;

-- name: SetEventPaid :exec
UPDATE events
SET
    paid = $2,
    updated_at = NOW()
WHERE
    id = $1;

-- name: DeleteEvents :exec
DELETE FROM events
WHERE id = ANY($1::uuid[]);