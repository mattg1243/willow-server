package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/mattg1243/willow-server/db"
	custom_middleware "github.com/mattg1243/willow-server/middleware"
)

// CreatePaymentTypeHandler Create a custom payment type
func (h *Handler) CreatePaymentTypeHandler(w http.ResponseWriter, r *http.Request) {
	req := &createPaymentTypeRequest{}

	if err := req.bind(r, h.validator); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user := custom_middleware.GetUserFromContext(r)
	userID, err := uuid.Parse(user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newPaymentType, err := h.queries.CreatePaymentType(r.Context(), db.CreatePaymentTypeParams{
		UserID: pgtype.UUID{Bytes: userID, Valid: true},
		Name:   req.Name,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(newPaymentType); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// GetPaymentTypesHandler get all default payment types and users custom types
// or if ?id is provided, get a specific payment type
func (h *Handler) GetPaymentTypesHandler(w http.ResponseWriter, r *http.Request) {
	user := custom_middleware.GetUserFromContext(r)
	userID, err := uuid.Parse(user)
	paymentTypeIDQuery := r.URL.Query().Get("id")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if paymentTypeIDQuery != "" {
		// Get specific payment type
		paymentTypeID64, err := strconv.ParseInt(paymentTypeIDQuery, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		paymentTypeID := int32(paymentTypeID64)

		paymentType, err := h.queries.GetPaymentType(r.Context(), db.GetPaymentTypeParams{
			ID:     paymentTypeID,
			UserID: pgtype.UUID{Bytes: userID, Valid: true},
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode(paymentType); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	} else { // Get all payment types
		// Stores all payment types, both default and custom
		var paymentTypes []db.PaymentType

		// Default payment types have a "null" user_id value, so we pass in a null UUID
		// to the query to get the default payment types.
		paymentTypesDefault, err := h.queries.GetPaymentTypes(r.Context(), pgtype.UUID{Valid: false})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Fetch custom payment types for the user
		paymentTypesCustom, err := h.queries.GetPaymentTypes(r.Context(), pgtype.UUID{Bytes: userID, Valid: true})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Combine default and custom payment types
		paymentTypes = append(paymentTypes, paymentTypesDefault...)
		paymentTypes = append(paymentTypes, paymentTypesCustom...)

		if err := json.NewEncoder(w).Encode(paymentTypes); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

// UpdatePaymentTypeHandler update users own custom payment type, NOT default types
func (h *Handler) UpdatePaymentTypeHandler(w http.ResponseWriter, r *http.Request) {
	req := &updatePaymentTypeRequest{}

	if err := req.bind(r, h.validator); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Ownership check: Ensure the user in the request context matches the user ID in the request body
	if custom_middleware.GetUserFromContext(r) != req.UserID {
		http.Error(w, "Unauthorized user", http.StatusUnauthorized)
		return
	}

	userID, err := uuid.Parse(req.UserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	paymentTypeID64, err := strconv.ParseInt(req.ID, 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	paymentTypeID32 := int32(paymentTypeID64)

	paymentType, err := h.queries.UpdatePaymentType(r.Context(), db.UpdatePaymentTypeParams{
		UserID: pgtype.UUID{Bytes: userID, Valid: true},
		ID:     paymentTypeID32,
		Name:   req.Name,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(paymentType); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DeletePaymentTypeHandler delete users own custom payment, NOT default types
func (h *Handler) DeletePaymentTypeHandler(w http.ResponseWriter, r *http.Request) {
	req := &deletePaymentTypeRequest{}

	if err := req.bind(r, h.validator); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Ownership check: Ensure the user in the request context matches the user ID in the request body
	if custom_middleware.GetUserFromContext(r) != req.UserID {
		http.Error(w, "Unauthorized user", http.StatusUnauthorized)
		return
	}

	userID, err := uuid.Parse(req.UserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	paymentTypeID64, err := strconv.ParseInt(req.ID, 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	paymentTypeID32 := int32(paymentTypeID64)

	err = h.queries.DeletePaymentType(r.Context(), db.DeletePaymentTypeParams{
		UserID: pgtype.UUID{Bytes: userID, Valid: true},
		ID:     paymentTypeID32,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}
