// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: events.sql

package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const createEvent = `-- name: CreateEvent :one
INSERT INTO events (
    client_id, user_id, date, duration, event_type_id, detail, rate, amount, running_balance, id, created_at, updated_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, NOW(), NOW()
) RETURNING id, user_id, client_id, date, duration, event_type_id, detail, rate, amount, running_balance, paid, created_at, updated_at
`

type CreateEventParams struct {
	ClientID       uuid.UUID        `json:"client_id"`
	UserID         uuid.UUID        `json:"user_id"`
	Date           pgtype.Timestamp `json:"date"`
	Duration       pgtype.Numeric   `json:"duration"`
	EventTypeID    uuid.UUID        `json:"event_type_id"`
	Detail         pgtype.Text      `json:"detail"`
	Rate           int32            `json:"rate"`
	Amount         int32            `json:"amount"`
	RunningBalance int32            `json:"running_balance"`
	ID             uuid.UUID        `json:"id"`
}

func (q *Queries) CreateEvent(ctx context.Context, arg CreateEventParams) (Event, error) {
	row := q.db.QueryRow(ctx, createEvent,
		arg.ClientID,
		arg.UserID,
		arg.Date,
		arg.Duration,
		arg.EventTypeID,
		arg.Detail,
		arg.Rate,
		arg.Amount,
		arg.RunningBalance,
		arg.ID,
	)
	var i Event
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.ClientID,
		&i.Date,
		&i.Duration,
		&i.EventTypeID,
		&i.Detail,
		&i.Rate,
		&i.Amount,
		&i.RunningBalance,
		&i.Paid,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteEvents = `-- name: DeleteEvents :exec
DELETE FROM events
WHERE id = ANY($1::uuid[])
`

func (q *Queries) DeleteEvents(ctx context.Context, dollar_1 []uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteEvents, dollar_1)
	return err
}

const getEvent = `-- name: GetEvent :one
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
WHERE e.id = $1
`

type GetEventRow struct {
	ID             uuid.UUID          `json:"id"`
	ClientID       uuid.UUID          `json:"client_id"`
	UserID         uuid.UUID          `json:"user_id"`
	Date           pgtype.Timestamptz `json:"date"`
	Duration       pgtype.Numeric     `json:"duration"`
	EventTypeID    uuid.UUID          `json:"event_type_id"`
	EventTypeTitle string             `json:"event_type_title"`
	Detail         pgtype.Text        `json:"detail"`
	Rate           int32              `json:"rate"`
	Amount         int32              `json:"amount"`
	RunningBalance int32              `json:"running_balance"`
	Paid           pgtype.Bool        `json:"paid"`
	Charge         bool               `json:"charge"`
}

func (q *Queries) GetEvent(ctx context.Context, id uuid.UUID) (GetEventRow, error) {
	row := q.db.QueryRow(ctx, getEvent, id)
	var i GetEventRow
	err := row.Scan(
		&i.ID,
		&i.ClientID,
		&i.UserID,
		&i.Date,
		&i.Duration,
		&i.EventTypeID,
		&i.EventTypeTitle,
		&i.Detail,
		&i.Rate,
		&i.Amount,
		&i.RunningBalance,
		&i.Paid,
		&i.Charge,
	)
	return i, err
}

const getEvents = `-- name: GetEvents :many
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
ORDER BY e.date ASC
`

type GetEventsParams struct {
	ClientID uuid.UUID          `json:"client_id"`
	Column2  pgtype.Timestamptz `json:"column_2"`
	Column3  pgtype.Timestamptz `json:"column_3"`
}

type GetEventsRow struct {
	ID             uuid.UUID          `json:"id"`
	UserID         uuid.UUID          `json:"user_id"`
	ClientID       uuid.UUID          `json:"client_id"`
	Date           pgtype.Timestamptz `json:"date"`
	Duration       pgtype.Numeric     `json:"duration"`
	EventTypeID    uuid.UUID          `json:"event_type_id"`
	EventTypeTitle string             `json:"event_type_title"`
	Detail         pgtype.Text        `json:"detail"`
	Rate           int32              `json:"rate"`
	Amount         int32              `json:"amount"`
	RunningBalance int32              `json:"running_balance"`
	Paid           pgtype.Bool        `json:"paid"`
	Charge         bool               `json:"charge"`
}

func (q *Queries) GetEvents(ctx context.Context, arg GetEventsParams) ([]GetEventsRow, error) {
	rows, err := q.db.Query(ctx, getEvents, arg.ClientID, arg.Column2, arg.Column3)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetEventsRow
	for rows.Next() {
		var i GetEventsRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.ClientID,
			&i.Date,
			&i.Duration,
			&i.EventTypeID,
			&i.EventTypeTitle,
			&i.Detail,
			&i.Rate,
			&i.Amount,
			&i.RunningBalance,
			&i.Paid,
			&i.Charge,
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

const getPayoutEvents = `-- name: GetPayoutEvents :many
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
ORDER BY e.date ASC
`

type GetPayoutEventsRow struct {
	ID             uuid.UUID          `json:"id"`
	UserID         uuid.UUID          `json:"user_id"`
	ClientID       uuid.UUID          `json:"client_id"`
	Date           pgtype.Timestamptz `json:"date"`
	Duration       pgtype.Numeric     `json:"duration"`
	EventTypeID    uuid.UUID          `json:"event_type_id"`
	EventTypeTitle string             `json:"event_type_title"`
	Detail         pgtype.Text        `json:"detail"`
	Rate           int32              `json:"rate"`
	Amount         int32              `json:"amount"`
	RunningBalance int32              `json:"running_balance"`
	Paid           pgtype.Bool        `json:"paid"`
	Charge         bool               `json:"charge"`
}

func (q *Queries) GetPayoutEvents(ctx context.Context, payoutID uuid.UUID) ([]GetPayoutEventsRow, error) {
	rows, err := q.db.Query(ctx, getPayoutEvents, payoutID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetPayoutEventsRow
	for rows.Next() {
		var i GetPayoutEventsRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.ClientID,
			&i.Date,
			&i.Duration,
			&i.EventTypeID,
			&i.EventTypeTitle,
			&i.Detail,
			&i.Rate,
			&i.Amount,
			&i.RunningBalance,
			&i.Paid,
			&i.Charge,
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

const markEventPaid = `-- name: MarkEventPaid :one
UPDATE events
SET 
    paid = $2,
    updated_at = NOW()

WHERE
    id = $1
RETURNING id, user_id, client_id, date, duration, event_type_id, detail, rate, amount, running_balance, paid, created_at, updated_at
`

type MarkEventPaidParams struct {
	ID   uuid.UUID   `json:"id"`
	Paid pgtype.Bool `json:"paid"`
}

func (q *Queries) MarkEventPaid(ctx context.Context, arg MarkEventPaidParams) (Event, error) {
	row := q.db.QueryRow(ctx, markEventPaid, arg.ID, arg.Paid)
	var i Event
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.ClientID,
		&i.Date,
		&i.Duration,
		&i.EventTypeID,
		&i.Detail,
		&i.Rate,
		&i.Amount,
		&i.RunningBalance,
		&i.Paid,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const setEventPaid = `-- name: SetEventPaid :exec
UPDATE events
SET
    paid = $2,
    updated_at = NOW()
WHERE
    id = $1
`

type SetEventPaidParams struct {
	ID   uuid.UUID   `json:"id"`
	Paid pgtype.Bool `json:"paid"`
}

func (q *Queries) SetEventPaid(ctx context.Context, arg SetEventPaidParams) error {
	_, err := q.db.Exec(ctx, setEventPaid, arg.ID, arg.Paid)
	return err
}

const setEventsToMiscType = `-- name: SetEventsToMiscType :exec
UPDATE events
SET event_type_id = $2
WHERE event_type_id = $1
`

type SetEventsToMiscTypeParams struct {
	EventTypeID   uuid.UUID `json:"event_type_id"`
	EventTypeID_2 uuid.UUID `json:"event_type_id_2"`
}

func (q *Queries) SetEventsToMiscType(ctx context.Context, arg SetEventsToMiscTypeParams) error {
	_, err := q.db.Exec(ctx, setEventsToMiscType, arg.EventTypeID, arg.EventTypeID_2)
	return err
}

const updateEvent = `-- name: UpdateEvent :one
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
RETURNING id, user_id, client_id, date, duration, event_type_id, detail, rate, amount, running_balance, paid, created_at, updated_at
`

type UpdateEventParams struct {
	ID             uuid.UUID        `json:"id"`
	Date           pgtype.Timestamp `json:"date"`
	Duration       pgtype.Numeric   `json:"duration"`
	EventTypeID    uuid.UUID        `json:"event_type_id"`
	Detail         pgtype.Text      `json:"detail"`
	Rate           int32            `json:"rate"`
	Amount         int32            `json:"amount"`
	RunningBalance int32            `json:"running_balance"`
	Paid           pgtype.Bool      `json:"paid"`
}

func (q *Queries) UpdateEvent(ctx context.Context, arg UpdateEventParams) (Event, error) {
	row := q.db.QueryRow(ctx, updateEvent,
		arg.ID,
		arg.Date,
		arg.Duration,
		arg.EventTypeID,
		arg.Detail,
		arg.Rate,
		arg.Amount,
		arg.RunningBalance,
		arg.Paid,
	)
	var i Event
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.ClientID,
		&i.Date,
		&i.Duration,
		&i.EventTypeID,
		&i.Detail,
		&i.Rate,
		&i.Amount,
		&i.RunningBalance,
		&i.Paid,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateRunningBalance = `-- name: UpdateRunningBalance :exec
UPDATE events
SET
    running_balance = $2
WHERE
    id = $1
`

type UpdateRunningBalanceParams struct {
	ID             uuid.UUID `json:"id"`
	RunningBalance int32     `json:"running_balance"`
}

func (q *Queries) UpdateRunningBalance(ctx context.Context, arg UpdateRunningBalanceParams) error {
	_, err := q.db.Exec(ctx, updateRunningBalance, arg.ID, arg.RunningBalance)
	return err
}
