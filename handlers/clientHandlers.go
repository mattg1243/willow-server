package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/mattg1243/willow-server/db"
)

func (h *Handler) CreateClientHandler(c *fiber.Ctx) error {
	var client db.Client
	req := &createClientRequest{}

	if err := req.bind(c, &client, h.validator); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(err.Error())
	}

	user := c.Locals("user").(string)
	userID, err := uuid.Parse(user)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON("UUID conversion error")
	}
	newClient, err := h.queries.CreateClient(c.Context(), db.CreateClientParams{
		UserID: userID,
		Fname:  client.Fname,
		Lname:  client.Lname,
		Email:  client.Email,
		Phone: client.Phone,
		Rate: client.Rate,
		Balancenotifythreshold: client.Balancenotifythreshold,
		ID:     uuid.New(),
	})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(http.StatusCreated).JSON(newClient)
}

func (h *Handler) GetClientHandler(c *fiber.Ctx) error {
	user := c.Locals("user")
	userID, err := uuid.Parse(user.(string))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON("UUID conversion error")
	}

	clientIDStr := c.Queries()["id"]
	if clientIDStr == "" {
		clients, err := h.queries.GetClients(c.Context(), userID)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(err.Error())
		}
		return c.Status(200).JSON(clients)
	} else {
		id, err := uuid.Parse(clientIDStr)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(err.Error())
		}
		clients, err := h.queries.GetClient(c.Context(), id)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(err.Error())
		}
		return c.Status(200).JSON(clients)
	}
}

func (h *Handler) UpdateClientHandler(c *fiber.Ctx) error {
	clientIDStr := c.Queries()["id"]
	clientID, err := uuid.Parse(clientIDStr)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	var req updateClientRequest
	var client db.Client

	if err := req.bind(c, &client, h.validator); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{"error": err.Error()})
	}

	err = h.queries.UpdateClient(c.Context(), db.UpdateClientParams{
		ID:                     clientID,
		Fname:                  client.Fname,
		Lname:                  client.Lname,
		Email:                  client.Email,
		Phone: 									client.Phone,
		Balancenotifythreshold: client.Balancenotifythreshold,
		Rate:                   client.Rate,
		Isarchived:             client.Isarchived,
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

func (h *Handler) DeleteClientHandler(c *fiber.Ctx) error {
	clientIDStr := c.Queries()["id"]
	clientID, err := uuid.Parse(clientIDStr)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err = h.queries.DeleteClient(c.Context(), clientID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.SendStatus(200)
}