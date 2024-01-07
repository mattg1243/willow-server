package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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

	userId, err := uuid.Parse(claims.id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON("UUID conversion error")
	}
	newClient, err := h.queries.CreateClient(c.Context(), db.CreateClientParams{UserID: userId })

	return c.Status(http.StatusCreated).JSON(newClient)
}

func (h *Handler) GetClientHandler(c* fiber.Ctx) error {
	clientIdStr := c.Params("id")
	clientId, err := uuid.Parse(clientIdStr)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	client, err := h.queries.GetClient(c.Context(), clientId)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(200).JSON(client)
}