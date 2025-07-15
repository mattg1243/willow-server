package handlers

import (
	"net/http"
)

// save custom payment type to db
func (h *Handler) CreatePaymentTypeHandler(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusCreated)
}

// gets single payment type by id
func (h *Handler) GetPaymentTypeHandler(w http.ResponseWriter, r *http.Request) {

}

// get all default payment types and users custom types
func (h *Handler) GetPaymentTypesHandler(w http.ResponseWriter, r *http.Request) {


}

// update users own custom payment type, NOT default types
func (h *Handler) UpdatePaymentTypeHandler(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
}

// delete users own custom payment, NOT default types
func (h *Handler) DeletePaymentTypeHandler(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
}
