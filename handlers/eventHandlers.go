package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/mattg1243/sqlc-fiber/db"
)

func (h *Handler) CreateEventHandler(c *fiber.Ctx) error {
	req := &createEventRequest{}
	event := db.Event{}

	if err := req.bind(c, &event, h.validator); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	newEvent, err := h.queries.CreateEvent(c.Context(), db.CreateEventParams{
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
