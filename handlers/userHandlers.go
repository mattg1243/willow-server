package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/mattg1243/willow-server/db"
	custom_middleware "github.com/mattg1243/willow-server/middleware"
	"github.com/mattg1243/willow-server/utils"
)

func (h *Handler) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	userIdStr := custom_middleware.GetUserFromContext(r)
	userId, err := uuid.Parse(userIdStr)

	if err != nil {
		http.Error(w, "User not found with request", http.StatusUnauthorized)
		return
	}
	// Get user from db
	user, err := h.queries.GetUser(r.Context(), userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	// Encode user as json and send to client
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Failed to encode user in the response", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetUserContactInfo(w http.ResponseWriter, r *http.Request) {
	userIdStr := custom_middleware.GetUserFromContext(r)
	userId, err := uuid.Parse(userIdStr)

	if err != nil {
		http.Error(w, "User not found with request", http.StatusUnauthorized)
		return
	}

	contactInfo, err := h.queries.GetUserContactInfo(r.Context(), userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(contactInfo); err != nil {
		http.Error(w, "Failed to encode contact info in the response", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user db.User
	var contactInfo db.UserContactInfo
	req := &createUserRequest{}

	if err := req.bind(r, &user, &contactInfo, h.validator); err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	// hash the password
	hash, err := user.HashPassword(req.User.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// save user
	newUser, err := h.queries.CreateUser(r.Context(), db.CreateUserParams{
		Hash:          hash,
		Email:         user.Email,
		Fname:         user.Fname,
		Lname:         user.Lname,
		Rate:          user.Rate,
		Nameforheader: user.Nameforheader,
		ID:            uuid.New(),
	})

	if err != nil {
		// Check if it's a Postgres error
		if pgErr, ok := err.(*pgconn.PgError); ok {
			// Check if it's a unique violation error
			if pgErr.Code == pgerrcode.UniqueViolation {
				http.Error(w, "User already exists", http.StatusBadRequest)
				return
			}
		}
		// Return a generic error for any other cases
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// save contact info
	_, err = h.queries.CreateUserContactInfo(r.Context(), db.CreateUserContactInfoParams{
		ID:     uuid.New(),
		Phone:  contactInfo.Phone,
		City:   contactInfo.City,
		State:  contactInfo.State,
		Street: contactInfo.Street,
		Zip:    contactInfo.Zip,
		UserID: newUser.ID,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(newUser); err != nil {
		http.Error(w, "Failed to encode user in the response", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user db.User
	var contactInfo db.UserContactInfo
	req := &updateUserRequest{}

	// parse user id from claims
	userIdStr := custom_middleware.GetUserFromContext(r)
	userId, err := uuid.Parse(userIdStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := req.bind(r, &user, &contactInfo, h.validator); err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	// save user
	updatedUser, err := h.queries.UpdateUser(r.Context(), db.UpdateUserParams{
		Fname:         user.Fname,
		Lname:         user.Lname,
		Nameforheader: user.Nameforheader,
		License:       user.License,
		Rate:          user.Rate,
		ID:            userId,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// save contact info
	_, err = h.queries.UpdateUserContactInfo(r.Context(), db.UpdateUserContactInfoParams{
		UserID:      userId,
		Phone:       contactInfo.Phone,
		City:        contactInfo.City,
		State:       contactInfo.State,
		Street:      contactInfo.Street,
		Zip:         contactInfo.Zip,
		Paymentinfo: contactInfo.Paymentinfo,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(updatedUser); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	userIdStr := custom_middleware.GetUserFromContext(r)
	userId, err := uuid.Parse(userIdStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.queries.DeleteUser(r.Context(), userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("User deleted successfully"))
}

func (h *Handler) LoginUserHandler(w http.ResponseWriter, r *http.Request) {

	req := loginUserRequest{}

	if err := req.bind(r, h.validator); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := h.queries.GetUserByEmail(r.Context(), req.Email)
	if err != nil {
		if err.Error() == "no rows in result set" {
			http.Error(w, "No user found with that email address", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	match := user.CheckPassword(req.Password)

	if match {
		payload := utils.JwtPayload{Id: user.ID.String(), Email: user.Email}
		jwt, err := utils.GenerateJWT(payload)
		if err != nil {
			fmt.Printf("%q", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Create and set cookie
		cookie := &http.Cookie{
			Name:     "willow-access-token",
			Value:    jwt,
			Expires:  time.Now().Add((time.Hour * 72)),
			HttpOnly: os.Getenv("PROD") == "true", //TODO change to true for production
			Secure:   os.Getenv("PROD") == "true", //TODO change to true for production
			Path:     "/",
			SameSite: http.SameSiteLaxMode,
		}

		http.SetCookie(w, cookie)
		w.WriteHeader(http.StatusOK)
		return

	} else {
		http.Error(w, "Invalid login credentials.", http.StatusUnauthorized)
		return
	}
}

func (h *Handler) LogoutUserHandler(w http.ResponseWriter, r *http.Request) {

	cookie := &http.Cookie{
		Name:     "willow-access-token",
		Value:    "",
		Expires:  time.Now().Add(-1 * time.Hour),
		HttpOnly: os.Getenv("PROD") == "true", // TODO: Set to true in production
		Secure:   os.Getenv("PROD") == "true", // TODO: Set to true in production
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, cookie)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Successfully logged out"))
}
