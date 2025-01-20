-- name: CreateEventType :one
INSERT INTO event_types (
  id, title, user_id, charge, created_at, updated_at
) values (
  $1, $2, $3, $4, NOW(), NOW()
) RETURNING *;

-- name: GetEventType :one
SELECT *
FROM event_types
WHERE id = $1;

-- name: GetEventTypes :many
SELECT *
FROM event_types
WHERE user_id = $1 OR user_id IS NULL;

-- name: UpdateEventType :one
UPDATE event_types
SET
  title = $2,
  charge = $3,
  updated_at = NOW()
WHERE
  id = $1
RETURNING *;

-- name: DeleteEventTypes :exec
DELETE FROM event_types
WHERE id = $1;