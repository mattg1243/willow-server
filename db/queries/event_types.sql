-- name: CreateEventType :one
INSERT INTO event_types (
  id, "name", user_id, charge, source
) values (
  $1, $2, $3, $4, "custom"
) RETURNING *;

-- name: GetEventType :one
SELECT *
FROM event_types
WHERE id = $1;

-- name: GetEventTypes :many
SELECT *
FROM event_types
WHERE user_id = $1;

-- name: UpdateEventType :exec
UPDATE event_types
SET
  "name" = $2,
  charge = $3
WHERE
  id = $1;

-- name: DeleteEventTypes :exec
DELETE FROM event_types
WHERE id = $1;