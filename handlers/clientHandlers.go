package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/mattg1243/willow-server/db"
	custom_middleware "github.com/mattg1243/willow-server/middleware"
)

func (h *Handler) CreateClientHandler(w http.ResponseWriter, r *http.Request) {
	var client db.Client
	req := &createClientRequest{}

	if err := req.bind(r, &client, h.validator); err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	userIDStr := custom_middleware.GetUserFromContext(r)
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	newClient, err := h.queries.CreateClient(r.Context(), db.CreateClientParams{
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(newClient); err != nil {
		http.Error(w, "Failed to encode client in the response", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetClientHandler(w http.ResponseWriter, r *http.Request)  {
	userIDStr := custom_middleware.GetUserFromContext(r)
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	clientIDStr := r.URL.Query().Get("id")
	if clientIDStr == "" {
		clients, err := h.queries.GetClients(r.Context(), userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(clients); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		id, err := uuid.Parse(clientIDStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		client, err := h.queries.GetClient(r.Context(), id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(client); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (h *Handler) UpdateClientHandler(w http.ResponseWriter, r *http.Request) {
	clientIDStr := r.URL.Query().Get("id")
	clientID, err := uuid.Parse(clientIDStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var req updateClientRequest
	var client db.Client

	if err := req.bind(r, &client, h.validator); err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	err = h.queries.UpdateClient(r.Context(), db.UpdateClientParams{
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	updatedClient, err := h.queries.GetClient(r.Context(), clientID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(updatedClient); err != nil {
		http.Error(w, "Failed to encode updated client in response", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) DeleteClientHandler(w http.ResponseWriter, r *http.Request) {
	clientIDStr := r.URL.Query()["id"][0]
	clientID, err := uuid.Parse(clientIDStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.queries.DeleteClient(r.Context(), clientID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Client deleted successfully"))
}