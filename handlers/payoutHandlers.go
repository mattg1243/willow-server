package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/mattg1243/willow-server/core"
	"github.com/mattg1243/willow-server/db"
	custom_middleware "github.com/mattg1243/willow-server/middleware"
)


type payoutStruct struct {
	Payout 		int							`json:"payout"`
	Events		[]uuid.UUID		`json:"events"`
}

func (h *Handler) MakePayoutHandler(w http.ResponseWriter, r *http.Request) {
	userIdStr := custom_middleware.GetUserFromContext(r)
	userId, err := uuid.Parse(userIdStr)
	
	if err != nil {
		http.Error(w, "User not found with request", http.StatusUnauthorized)
		return
	}

	payout := 0
	var events []db.GetEventsRow

	// Check if for specific client or global
	clientIdStr := r.URL.Query().Get("client")
	
	if clientIdStr != "" {
		clientId, err := uuid.Parse(clientIdStr)

		if err != nil {
			http.Error(w, "Invalid clientId provided with request", http.StatusUnauthorized)
			return
		}
		events, err = h.queries.GetEvents(r.Context(), clientId)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	} else {
		events, err = h.queries.GetEvents(r.Context(), userId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	payout, chargedEvents := core.CalculatePayout(events)

	res := payoutStruct{ Payout: payout, Events: chargedEvents }

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) SavePayoutHandler(w http.ResponseWriter, r *http.Request) {
	userIdStr := custom_middleware.GetUserFromContext(r)
	userId, err := uuid.Parse(userIdStr)
	
	if err != nil {
		http.Error(w, "User not found with request", http.StatusUnauthorized)
		return
	}

	req := payoutRequest{}
	err = req.bind(r, h.validator) 
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	payoutID := uuid.New()

	params := db.CreatePayoutParams{
		ID: payoutID,
		UserID: userId,
		ClientID: pgtype.UUID{Bytes: req.ClientID, Valid: true},
		Amount: int32(req.Payout),
	}
	// TODO make this process a tx
	// Save payout
	_, err = h.queries.CreatePayout(r.Context(), params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	// Add events to join table
	for i := 0; i < len(req.Events); i++ {
		// Add event to join table
		_, err = h.queries.AddEventToPayout(r.Context(), db.AddEventToPayoutParams{
			PayoutID: payoutID,
			EventID: req.Events[i],
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Mark event as paid
		_, err = h.queries.MarkEventPaid(
			r.Context(), 
			db.MarkEventPaidParams{ 
				ID: req.Events[i],
				Paid: pgtype.Bool{ 
					Bool: true, 
					Valid: true,
				},
			})
			if err != nil {
				fmt.Print("error marking event as paid")
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) GetPayoutHandler(w http.ResponseWriter, r *http.Request) {
	payoutIDStr := r.URL.Query().Get("id")
	if payoutIDStr == "" {
		// Return all of the users payouts
		userIDStr := custom_middleware.GetUserFromContext(r)
		if userIDStr == "" {
			http.Error(w, "No user ID provided with request", http.StatusBadRequest)
			return
		}

		userID, err := uuid.Parse(userIDStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		payouts, err := h.queries.GetPayouts(r.Context(), userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(payouts); err != nil {
			http.Error(w, "Failed to encode payout data in the response", http.StatusInternalServerError)
			return
		}
	} else {
		// Return only the specified payout
		payoutID, err := uuid.Parse(payoutIDStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		payout, err := h.queries.GetPayout(r.Context(), payoutID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(payout); err != nil {
			http.Error(w, "Failed to encode payout data in the response", http.StatusInternalServerError)
			return
		}
	}
}

func (h *Handler) GetPayoutsHandler(w http.ResponseWriter, r *http.Request) {
	userIDStr := custom_middleware.GetUserFromContext(r)
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	payouts, err := h.queries.GetPayouts(r.Context(), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(payouts); err != nil {
		http.Error(w, "Failed to encode payout data in the response", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) DeletePayoutHandler(w http.ResponseWriter, r *http.Request) {
	// Get all events included in payout
	payoutIDStr := r.URL.Query().Get("id")
	if payoutIDStr == "" {
		http.Error(w, "No payout ID provided with request", http.StatusBadRequest)
		return
	}

	payoutID, err := uuid.Parse(payoutIDStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	events, err := h.queries.GetPayoutEvents(r.Context(), payoutID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Mark events as unpaid
	for i := 0; i < len(events); i++ {
		_, err := h.queries.MarkEventPaid(
			r.Context(), 
			db.MarkEventPaidParams{ 
				ID: events[i].ID, 
				Paid: pgtype.Bool{ Bool: false, Valid: true },
			})

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
	}
	// Delete payout row
	err = h.queries.DeletePayout(r.Context(), payoutID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}