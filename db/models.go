// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Client struct {
	ID                     uuid.UUID        `json:"id"`
	UserID                 uuid.UUID        `json:"user_id"`
	Fname                  string           `json:"fname"`
	Lname                  pgtype.Text      `json:"lname"`
	Email                  pgtype.Text      `json:"email"`
	Phone                  pgtype.Text      `json:"phone"`
	Balance                int32            `json:"balance"`
	Balancenotifythreshold int32            `json:"balancenotifythreshold"`
	Rate                   int32            `json:"rate"`
	Isarchived             pgtype.Bool      `json:"isarchived"`
	CreatedAt              pgtype.Timestamp `json:"created_at"`
	UpdatedAt              pgtype.Timestamp `json:"updated_at"`
}

type Event struct {
	ID             uuid.UUID        `json:"id"`
	UserID         uuid.UUID        `json:"user_id"`
	ClientID       uuid.UUID        `json:"client_id"`
	Date           pgtype.Timestamp `json:"date"`
	Duration       pgtype.Numeric   `json:"duration"`
	EventTypeID    uuid.UUID        `json:"event_type_id"`
	Detail         pgtype.Text      `json:"detail"`
	Rate           int32            `json:"rate"`
	Amount         int32            `json:"amount"`
	RunningBalance int32            `json:"running_balance"`
}

type EventType struct {
	ID        uuid.UUID        `json:"id"`
	UserID    pgtype.UUID      `json:"user_id"`
	Source    pgtype.Text      `json:"source"`
	Name      string           `json:"name"`
	Charge    bool             `json:"charge"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
}

type User struct {
	ID            uuid.UUID        `json:"id"`
	Fname         string           `json:"fname"`
	Lname         string           `json:"lname"`
	Email         string           `json:"email"`
	Hash          string           `json:"hash"`
	Nameforheader string           `json:"nameforheader"`
	License       pgtype.Text      `json:"license"`
	CreatedAt     pgtype.Timestamp `json:"created_at"`
	UpdatedAt     pgtype.Timestamp `json:"updated_at"`
}

type UserContactInfo struct {
	ID          uuid.UUID        `json:"id"`
	UserID      uuid.UUID        `json:"user_id"`
	Phone       pgtype.Text      `json:"phone"`
	City        pgtype.Text      `json:"city"`
	State       pgtype.Text      `json:"state"`
	Street      pgtype.Text      `json:"street"`
	Zip         pgtype.Text      `json:"zip"`
	Paymentinfo []byte           `json:"paymentinfo"`
	CreatedAt   pgtype.Timestamp `json:"created_at"`
	UpdatedAt   pgtype.Timestamp `json:"updated_at"`
}
