package handlers

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mattg1243/willow-server/db"
)

type Handler struct {
	dbPool      *pgxpool.Pool
	queries   *db.Queries
	validator *Validator
}

func New(dbPool *pgxpool.Pool) *Handler {
	v := NewValidator()
	return &Handler{dbPool, db.New(dbPool), v}
}