package custom_middleware

import (
	"context"
	"net/http"

	"github.com/mattg1243/willow-server/utils"
)

type contextKey string

const (
	UserIDContextKey contextKey = "user"
	EmailContextKey contextKey = "email"
)

func AuthJwt(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		cookie, err := r.Cookie("willow-access-token")
		if err != nil || cookie.Value == "" {
			http.Error(w, "Unauthorized: Missing or invalid cookie", http.StatusUnauthorized)
			return
		}

		token := cookie.Value

		claims, err := utils.ValidateJWT(token)
		if err != nil {
			http.Error(w, "Unauthorized: Missing or invalid cookie", http.StatusUnauthorized)
			return
		}

		  // Create a new context with the user and email values
			ctx := context.WithValue(r.Context(), UserIDContextKey, claims.Id)
			ctx = context.WithValue(ctx, EmailContextKey, claims.Email)

			// Call the next handler with the new context
			next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// TODO create middleware that verifies user ownership of events and clients