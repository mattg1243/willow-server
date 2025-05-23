// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: event_types.sql

package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const createEventType = `-- name: CreateEventType :one
INSERT INTO event_types (
  id, title, user_id, charge, created_at, updated_at
) values (
  $1, $2, $3, $4, NOW(), NOW()
) RETURNING id, user_id, source, title, charge, created_at, updated_at
`

type CreateEventTypeParams struct {
	ID     uuid.UUID   `json:"id"`
	Title  string      `json:"title"`
	UserID pgtype.UUID `json:"user_id"`
	Charge bool        `json:"charge"`
}

func (q *Queries) CreateEventType(ctx context.Context, arg CreateEventTypeParams) (EventType, error) {
	row := q.db.QueryRow(ctx, createEventType,
		arg.ID,
		arg.Title,
		arg.UserID,
		arg.Charge,
	)
	var i EventType
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Source,
		&i.Title,
		&i.Charge,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteEventTypes = `-- name: DeleteEventTypes :exec
DELETE FROM event_types
WHERE id = $1
`

func (q *Queries) DeleteEventTypes(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteEventTypes, id)
	return err
}

const getEventType = `-- name: GetEventType :one
SELECT id, user_id, source, title, charge, created_at, updated_at
FROM event_types
WHERE id = $1
`

func (q *Queries) GetEventType(ctx context.Context, id uuid.UUID) (EventType, error) {
	row := q.db.QueryRow(ctx, getEventType, id)
	var i EventType
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Source,
		&i.Title,
		&i.Charge,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getEventTypes = `-- name: GetEventTypes :many
SELECT id, user_id, source, title, charge, created_at, updated_at
FROM event_types
WHERE user_id = $1 OR user_id IS NULL
`

func (q *Queries) GetEventTypes(ctx context.Context, userID pgtype.UUID) ([]EventType, error) {
	rows, err := q.db.Query(ctx, getEventTypes, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []EventType
	for rows.Next() {
		var i EventType
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Source,
			&i.Title,
			&i.Charge,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateEventType = `-- name: UpdateEventType :one
UPDATE event_types
SET
  title = $2,
  charge = $3,
  updated_at = NOW()
WHERE
  id = $1
RETURNING id, user_id, source, title, charge, created_at, updated_at
`

type UpdateEventTypeParams struct {
	ID     uuid.UUID `json:"id"`
	Title  string    `json:"title"`
	Charge bool      `json:"charge"`
}

func (q *Queries) UpdateEventType(ctx context.Context, arg UpdateEventTypeParams) (EventType, error) {
	row := q.db.QueryRow(ctx, updateEventType, arg.ID, arg.Title, arg.Charge)
	var i EventType
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Source,
		&i.Title,
		&i.Charge,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
