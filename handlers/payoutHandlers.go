package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/mattg1243/willow-server/core"
	"github.com/mattg1243/willow-server/db"
	custom_middleware "github.com/mattg1243/willow-server/middleware"
)


type getPayoutResponse struct {
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

	res := getPayoutResponse{ Payout: payout, Events: chargedEvents }

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) SavePayoutHandler(w http.ResponseWriter, r *http.Request) {
	
}

func (h *Handler) GetPayoutHandler(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) UndoPayoutHandler(w http.ResponseWriter, r *http.Request) {
	// mark included events as paid
}

func (h *Handler) RedoPayoutHandler(w http.ResponseWriter, r *http.Request) {
	
}