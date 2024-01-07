package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/mattg1243/sqlc-fiber/db"
)

type JwtClaims struct {
	id    string
	email string
}

func (h *Handler) CreateClientHandler(c *fiber.Ctx) error {
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

	userID, err := uuid.Parse(claims.id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON("UUID conversion error")
	}
	newClient, err := h.queries.CreateClient(c.Context(), db.CreateClientParams{UserID: userID})

	return c.Status(http.StatusCreated).JSON(newClient)
}

func (h *Handler) GetClientHandler(c *fiber.Ctx) error {
	clientIDStr := c.Params("id")
	clientID, err := uuid.Parse(clientIDStr)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	client, err := h.queries.GetClient(c.Context(), clientID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(200).JSON(client)
}

func (h *Handler) GetClientsHandler(c *fiber.Ctx) error {
	user := c.Locals("user")
	claims, ok := user.(*JwtClaims)
	if !ok {
		return c.Status(http.StatusBadRequest).JSON("No user Id found with request")
	}

	userID, err := uuid.Parse(claims.id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON("UUID conversion error")
	}

	clients, err := h.queries.GetClients(c.Context(), userID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(200).JSON(clients)
}

func (h *Handler) UpdateClientHandler(c *fiber.Ctx) error {
	clientIDStr := c.Params("id")
	clientID, err := uuid.Parse(clientIDStr)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	var req updateClientRequest
	var model db.Client

	if err := req.bind(c, &model, h.validator); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{"error": err.Error()})
	}

	err = h.queries.UpdateClient(c.Context(), db.UpdateClientParams{
		ID:                     clientID,
		Fname:                  req.Client.Fname,
		Lname:                  pgtype.Text{String: req.Client.Lname},
		Email:                  pgtype.Text{String: req.Client.Email},
		Balance:                req.Client.Balance,
		Balancenotifythreshold: req.Client.Balancenotifythreshold,
		Rate:                   req.Client.Rate,
		Isarchived:             pgtype.Bool{Bool: req.Client.Isarchived},
	})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	updatedClient, err := h.queries.GetClient(c.Context(), clientID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve updated client"})
	}

	return c.Status(http.StatusOK).JSON(updatedClient)
}
