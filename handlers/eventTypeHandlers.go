package handlers

// import (
// 	"net/http"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/google/uuid"
// 	"github.com/mattg1243/willow-server/db"
// )

// // TODO make sure that when an event type is deleted, any event that references
// // that custom type has its type set the `Misc.` default type
// // TODO check that names dont overlap with default names

// func (h *Handler) CreateEventTypeHandler(c *fiber.Ctx) error {
// 	req := &createEventTypeRequest{}
// 	eventType := db.EventType{}

// 	if err := req.bind(c, &eventType, h.validator); err != nil {
// 		return c.Status(http.StatusBadRequest).JSON(err.Error())
// 	}

// 	user := c.Locals("user").(string)
// 	userID, err := uuid.Parse(user)

// 	if err != nil {
// 		return c.Status(http.StatusInternalServerError).JSON("UUID conversion error")
// 	}

// 	newEventType, err := h.queries.CreateEventType(c.Context(), db.CreateEventTypeParams{
// 		ID: uuid.New(),
// 		Name: eventType.Name,
// 		UserID: userID,
// 		Charge: eventType.Charge,
// 	})

// 	if err != nil {
// 		return c.Status(http.StatusInternalServerError).JSON(err.Error())
// 	}

// 	return c.Status(http.StatusCreated).JSON(newEventType)
// }