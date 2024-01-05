package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/mattg1243/sqlc-fiber/db"
)

type JwtClaims struct {
	id string
	email string
}

func (h *Handler) CreateClientHandler(c* fiber.Ctx) error {
	var client db.Client
	req := &createClientRequest{}

	if err := req.bind(c, &client, h.validator); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(err.Error())
	}

	user := c.Locals("user")
	claims, ok := user.(*JwtClaims)
	if !ok {
		return c.Status(http.StatusBadRequest).JSON("No user Id found with request")
	}
	
	var uuid pgtype.UUID
	// if err != nil {
	// 	return c.Status(http.StatusInternalServerError).JSON("UUID conversion error")
	// }
	// newClient, err := h.queries.CreateClient(c.Context(), db.CreateClientParams{UserID: uuid })
}