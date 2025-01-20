package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/mattg1243/willow-server/db"
	custom_middleware "github.com/mattg1243/willow-server/middleware"
)

// TODO make sure that when an event type is deleted, any event that references
// that custom type has its type set the `Misc.` default type
// TODO check that names dont overlap with default names

func (h *Handler) CreateEventTypeHandler(w http.ResponseWriter, r *http.Request) {
	req := &createEventTypeRequest{}
	eventType := db.EventType{}

	if err := req.bind(r, &eventType, h.validator); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user := custom_middleware.GetUserFromContext(r)
	userID, err := uuid.Parse(user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newEventType, err := h.queries.CreateEventType(r.Context(), db.CreateEventTypeParams{
		ID:     uuid.New(),
		Title:  eventType.Title,
		UserID: pgtype.UUID{Bytes: userID, Valid: true},
		Charge: eventType.Charge,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(newEventType); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetEventTypeHandler(w http.ResponseWriter, r *http.Request) {
	eventTypeIDQuery := r.URL.Query().Get("id")
	userIdStr := custom_middleware.GetUserFromContext(r)

	if eventTypeIDQuery != "" {
		// Get specific event type
		eventTypeID, err := uuid.Parse(eventTypeIDQuery)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		eventType, err := h.queries.GetEventType(r.Context(), eventTypeID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode(eventType); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		// Get all users event types
		userID, err := uuid.Parse(userIdStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		eventTypes, err := h.queries.GetEventTypes(r.Context(), pgtype.UUID{Bytes: userID, Valid: true})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode(eventTypes); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (h *Handler) UpdateEventTypeHandler(w http.ResponseWriter, r *http.Request) {
	var req updateEventTypeRequest
	var eventType db.EventType

	if err := req.bind(r, &eventType, h.validator); err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	updateEventType, err := h.queries.UpdateEventType(
		r.Context(),
		db.UpdateEventTypeParams{
			ID:     eventType.ID,
			Title:  eventType.Title,
			Charge: eventType.Charge,
		},
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(updateEventType); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) DeleteEventTypeHandler(w http.ResponseWriter, r *http.Request) {
	eventTypeIdQuery := r.URL.Query().Get("id")
	eventTypeID, err := uuid.Parse(eventTypeIdQuery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	miscID, err := uuid.FromBytes([]byte("3c8a8b7e-4e7d-4e8f-9a5b-6a4cf2e42b1c"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// Set all events that reference this type to Misc
	err = h.queries.SetEventsToMiscType(r.Context(), db.SetEventsToMiscTypeParams{
		EventTypeID:   eventTypeID,
		EventTypeID_2: miscID,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = h.queries.DeleteEventTypes(r.Context(), eventTypeID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}
