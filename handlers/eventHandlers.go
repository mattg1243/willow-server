package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/mattg1243/sqlc-fiber/db"
)

func (h *Handler) CreateEventHandler(c *fiber.Ctx) error {
	req := &createEventRequest{}
	event := db.Event{}

	if err := req.bind(c, &event, h.validator); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	newEvent, err := h.queries.CreateEvent(c.Context(), db.CreateEventParams{
		ClientID:   event.ClientID,
		Date:       event.Date,
		Duration:   event.Duration,
		Type:       event.Type,
		Detail:     event.Detail,
		Rate:       event.Rate,
		Amount:     event.Amount,
		Newbalance: event.Newbalance,
	})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(http.StatusCreated).JSON(newEvent)
}

func (h *Handler) GetEventHandler(c *fiber.Ctx) error {
	idStr := c.Query("id")

	if idStr == "" {
		return c.Status(http.StatusBadRequest).JSON("id is required")
	}

	id, err := uuid.Parse(idStr)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	event, err := h.queries.GetEvent(c.Context(), id)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(http.StatusFound).JSON(event)
}

func (h *Handler) GetEventsByClientHandler(c *fiber.Ctx) error {
	clientIdStr := c.Query("clientId")

	if clientIdStr == "" {
		return c.Status(http.StatusBadRequest).JSON("clientId is required")
	}

	clientId, err := uuid.Parse(clientIdStr)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	events, err := h.queries.GetClients(c.Context(), clientId)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(http.StatusFound).JSON(events)
}

func (h *Handler) UpdateEventHandler(c *fiber.Ctx) error {
	req := &updateEventRequest{}
	event := db.Event{}

	if err := req.bind(c, &event, h.validator); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	err := h.queries.UpdateEvent(c.Context(), db.UpdateEventParams{
		ID:         event.ID,
		Date:       event.Date,
		Duration:   event.Duration,
		Type:       event.Type,
		Detail:     event.Detail,
		Rate:       event.Rate,
		Amount:     event.Amount,
		Newbalance: event.Newbalance,
	})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return nil
}

func (h *Handler) DeleteEventHandler(c *fiber.Ctx) error {
	eventIdStr := c.Query("id")
	eventId, err := uuid.Parse(eventIdStr)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	queryErr := h.queries.DeleteEvent(c.Context(), eventId)
	if queryErr != nil {
		return c.Status(http.StatusInternalServerError).JSON(queryErr.Error())
	}

	return c.Status(http.StatusOK).JSON("event deleted")
}
