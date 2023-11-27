package handlers

import (
	"github.com/jackc/pgx/v5"
	"github.com/mattg1243/sqlc-fiber/db"
)

type Handler struct {
	Conn *pgx.Conn
	Queries *db.Queries
}

func NewHandler(conn *pgx.Conn) *Handler {
	return &Handler{conn, db.New(conn)}
}