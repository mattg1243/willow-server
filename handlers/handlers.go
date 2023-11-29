package handlers

import (
	"github.com/jackc/pgx/v5"
	"github.com/mattg1243/sqlc-fiber/db"
)

type Handler struct {
	conn *pgx.Conn
	queries *db.Queries
	validator *Validator
}

func NewHandler(conn *pgx.Conn) *Handler {
	v := NewValidator()
	return &Handler{conn, db.New(conn), v}
}