package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/mattg1243/willow-server/core"
	"github.com/mattg1243/willow-server/db"
	custom_middleware "github.com/mattg1243/willow-server/middleware"
)

func (h *Handler) CreateEventHandler(w http.ResponseWriter, r *http.Request) {
	req := &createEventRequest{}
	event := db.Event{}

	userIDStr := custom_middleware.GetUserFromContext(r)
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if err := req.bind(r, &event, h.validator); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newEvent, err := h.queries.CreateEvent(r.Context(), db.CreateEventParams{
		ClientID:       event.ClientID,
		UserID:         userID,
		ID:             uuid.New(),
		Date:           event.Date,
		Duration:       event.Duration,
		EventTypeID:    event.EventTypeID,
		Detail:         event.Detail,
		Rate:           event.Rate,
		Amount:         event.Amount,
		RunningBalance: 0,
	})

	if err != nil {
		log.Fatalf("error saving event: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// update balances
	allEvents, err := h.queries.GetEvents(r.Context(), db.GetEventsParams{ClientID: event.ClientID})
	if err != nil {
		log.Fatalf("error getting events")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	runningBalances := core.RecalcRunningBalances(allEvents)
	// save to db
	for i := 0; i < len(allEvents); i++ {
		err := h.queries.UpdateRunningBalance(
			r.Context(),
			db.UpdateRunningBalanceParams{ID: allEvents[i].ID, RunningBalance: int32(runningBalances[i])},
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	// update users balance to the value of the most recent events running balance value
	newClientBalance := runningBalances[len(runningBalances)-1]

	err = h.queries.UpdateClientBalance(
		r.Context(),
		db.UpdateClientBalanceParams{
			ID:      event.ClientID,
			Balance: int32(newClientBalance),
		},
	)

	if err := json.NewEncoder(w).Encode(newEvent); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetEventHandler(w http.ResponseWriter, r *http.Request) {
	eventIDQuery := r.URL.Query().Get("id")
	clientIDQuery := r.URL.Query().Get("clientId")
	payoutIDQuery := r.URL.Query().Get("payoutId")
	startDateQuery := r.URL.Query().Get("start")
	endDateQuery := r.URL.Query().Get("end")
	// Make sure one id param is provided
	if eventIDQuery == "" && clientIDQuery == "" && payoutIDQuery == "" {
		http.Error(w, "Neither clientId, eventId or payoutId provided with request", http.StatusInternalServerError)
		return
	}
	// Handle eventId
	if eventIDQuery != "" {
		eventID, err := uuid.Parse(eventIDQuery)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		event, err := h.queries.GetEvent(r.Context(), eventID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(event); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	// Handle clientId
	if clientIDQuery != "" {
		clientID, err := uuid.Parse(clientIDQuery)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Parse date range if provided
		var startDate *time.Time
		if startDateQuery != "" {
			startDateParsed, err := time.Parse(startDateQuery, time.RFC3339)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			startDate = &startDateParsed
		} else {
			startDate = nil
		}

		var endDate *time.Time
		if endDateQuery != "" {
			endDateParsed, err := time.Parse(endDateQuery, time.RFC3339)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			endDate = &endDateParsed
		} else {
			endDate = nil
		}

		// Construct params safely
		params := db.GetEventsParams{
			ClientID: clientID,
		}

		// Only assign values if dates are provided
		if startDate != nil {
			params.Column2 = pgtype.Timestamptz{Time: *startDate, Valid: true}
		}

		if endDate != nil {
			params.Column3 = pgtype.Timestamptz{Time: *endDate, Valid: true}
		}

		events, err := h.queries.GetEvents(r.Context(), params)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(events); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	// Handle payout id
	if payoutIDQuery != "" {
		payoutID, err := uuid.Parse(payoutIDQuery)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		events, err := h.queries.GetPayoutEvents(r.Context(), payoutID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(events); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (h *Handler) UpdateEventHandler(w http.ResponseWriter, r *http.Request) {
	var req updateEventRequest
	var event db.Event

	if err := req.bind(r, &event, h.validator); err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	updatedEvent, err := h.queries.UpdateEvent(r.Context(), db.UpdateEventParams{
		ID:          event.ID,
		Date:        event.Date,
		Duration:    event.Duration,
		EventTypeID: event.EventTypeID,
		Detail:      event.Detail,
		Rate:        event.Rate,
		Amount:      event.Amount,
		Paid:        event.Paid,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// update balances
	allEvents, err := h.queries.GetEvents(r.Context(), db.GetEventsParams{ClientID: event.ClientID})
	if err != nil {
		log.Fatalf("error getting events")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	runningBalances := core.RecalcRunningBalances(allEvents)
	// save to db
	for i := 0; i < len(allEvents); i++ {
		err := h.queries.UpdateRunningBalance(
			r.Context(),
			db.UpdateRunningBalanceParams{ID: allEvents[i].ID, RunningBalance: int32(runningBalances[i])},
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	// update client balance
	newClientBalance := runningBalances[len(runningBalances)-1]
	err = h.queries.UpdateClientBalance(
		r.Context(),
		db.UpdateClientBalanceParams{
			ID:      event.ClientID,
			Balance: int32(newClientBalance),
		},
	)

	if err := json.NewEncoder(w).Encode(updatedEvent); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (h *Handler) DeleteEventHandler(w http.ResponseWriter, r *http.Request) {
	eventIDsStr := r.URL.Query()["id"]
	clientIDStr := r.URL.Query().Get("client")
	userIDStr := custom_middleware.GetUserFromContext(r)

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if len(eventIDsStr) == 0 {
		http.Error(w, "No event id provided", http.StatusBadRequest)
		return
	}

	if len(clientIDStr) == 0 {
		http.Error(w, "No client id provided", http.StatusBadRequest)
		return
	}

	eventIDs := make([]uuid.UUID, len(eventIDsStr))
	for i, idStr := range eventIDsStr {
		id, err := uuid.Parse(idStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		eventIDs[i] = id
	}

	clientID, err := uuid.Parse(clientIDStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if event is associated with a payout
	inPayout, err := h.queries.EventIsInPayout(r.Context(), db.EventIsInPayoutParams{
		Column1: eventIDs,
		UserID:  userID,
	})

	fmt.Print("Event is in payout: ", inPayout)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if inPayout {
		http.Error(
			w,
			"One or more of the events you attempted to delete are associated with a payout. You must undo the associated payout before you can delete the event(s).",
			http.StatusBadRequest,
		)
		return
	}

	err = h.queries.DeleteEvents(r.Context(), eventIDs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// TODO implement this
	// update balances
	allEvents, err := h.queries.GetEvents(r.Context(), db.GetEventsParams{ClientID: clientID})
	if err != nil {
		log.Fatalf("error getting events")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	runningBalances := core.RecalcRunningBalances(allEvents)
	// save to db
	for i := 0; i < len(allEvents); i++ {
		err := h.queries.UpdateRunningBalance(
			r.Context(),
			db.UpdateRunningBalanceParams{ID: allEvents[i].ID, RunningBalance: int32(runningBalances[i])},
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	// update client balance
	var newBalance int
	if len(allEvents) == 0 {
		newBalance = 0
	} else {
		newBalance = runningBalances[len(runningBalances)-1]
	}
	err = h.queries.UpdateClientBalance(
		r.Context(),
		db.UpdateClientBalanceParams{
			ID:      clientID,
			Balance: int32(newBalance),
		},
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
