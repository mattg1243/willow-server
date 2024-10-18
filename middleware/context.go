package custom_middleware

import (
	"net/http"
)

// GetUserFromContext retrieves the user ID stored in the request context.
func GetUserFromContext(r *http.Request) string {
    if user, ok := r.Context().Value(UserIDContextKey).(string); ok {
        return user
    }
    return ""
}

// GetEmailFromContext retrieves the email stored in the request context.
func GetEmailFromContext(r *http.Request) string {
    if email, ok := r.Context().Value(EmailContextKey).(string); ok {
        return email
    }
    return ""
}