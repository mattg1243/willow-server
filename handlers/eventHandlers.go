package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/mattg1243/willow-server/db"
)

func (h *Handler) CreateEventHandler(c *fiber.Ctx) error {
	req := &createEventRequest{}
	event := db.Event{}

	if err := req.bind(c, &event, h.validator); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	// TODO calculate new balance

	newEvent, err := h.queries.CreateEvent(c.Context(), db.CreateEventParams{
		ClientID:   	event.ClientID,
		ID: 					uuid.New(),
		Date:       	event.Date,
		Duration:   	event.Duration,
		EventTypeID:	event.EventTypeID,
		Detail:     	event.Detail,
		Rate:       	event.Rate,
		Amount:     	event.Amount,
	})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(http.StatusCreated).JSON(newEvent)
}
