package handlers

import (
	"encoding/json"
	"fmt"
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
	var req updateClientRequest
	var client db.Client

	if err := req.bind(r, &client, h.validator); err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	fmt.Print()

	err := h.queries.UpdateClient(r.Context(), db.UpdateClientParams{
		ID:                     client.ID,
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

	updatedClient, err := h.queries.GetClient(r.Context(), client.ID)
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
	clientIDsStr := r.URL.Query()["id"]
	if len(clientIDsStr) == 0 {
		http.Error(w, "No ids provided", http.StatusBadRequest)
		return
	}

	clientIDs := make([]uuid.UUID, len(clientIDsStr))
	for i, idStr := range clientIDsStr {
		clientID, err := uuid.Parse(idStr)
		if err != nil {
			http.Error(w, "Invalid UUID: "+idStr, http.StatusBadRequest)
			return
		}
		clientIDs[i] = clientID
	}

	err := h.queries.DeleteClient(r.Context(), clientIDs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
 
	w.Write([]byte("Client(s) deleted successfully"))
}

func (h *Handler) BatchArchiveClientsHandler(w http.ResponseWriter, r *http.Request) {
	clientIDsStr := r.URL.Query()["id"]
	if len(clientIDsStr) == 0 {
		http.Error(w, "No ids provided", http.StatusBadRequest)
		return
	}

	clientIDs := make([]uuid.UUID, len(clientIDsStr))
	for i, idStr := range clientIDsStr {
		clientID, err := uuid.Parse(idStr)
		if err != nil {
			http.Error(w, "Invalid UUID: "+idStr, http.StatusBadRequest)
			return
		}
		clientIDs[i] = clientID
	}

	err := h.queries.BatchArchiveClients(r.Context(), clientIDs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Client(s) archived successfully"))
}