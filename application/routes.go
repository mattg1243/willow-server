package application

import (
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/mattg1243/willow-server/handlers"
	custom_middleware "github.com/mattg1243/willow-server/middleware"
)

func loadRoutes(h *handlers.Handler) *chi.Mux {
	// Create new chi router
	router := chi.NewRouter()
	// Global middleware
	attachMiddleware(router)
	// Attach base routes
	loadBaseRoutes(router, h)
	// Attach auth routes
	authRouter := chi.NewRouter()
	loadAuthRoutes(authRouter, h)
	router.Mount("/auth", authRouter)
	// Attach user routes
	userRouter := chi.NewRouter()
	loadUserRoutes(userRouter, h)
	router.Mount("/user", userRouter)
	// Attach client routes
	clientRouter := chi.NewRouter()
	loadClientRoutes(clientRouter, h)
	router.Mount("/client", clientRouter)
	// Attach event routes
	eventRouter := chi.NewRouter()
	loadEventRoutes(eventRouter, h)
	router.Mount("/event", eventRouter)
	// Attach event type routes
	eventTypeRouter := chi.NewRouter()
	loadEventTypeRoutes(eventTypeRouter, h)
	router.Mount("/event-types", eventTypeRouter)
	// Attach payout routes
	payoutRouter := chi.NewRouter()
	loadPayoutRoutes(payoutRouter, h)
	router.Mount("/payout", payoutRouter)
	// Attach payment type routes
	paymentTypeRouter := chi.NewRouter()
	loadPaymentTypeRoutes(paymentTypeRouter, h)
	router.Mount("/payment-types", paymentTypeRouter)
	// Return the completed router
	return router
}

func loadBaseRoutes(router chi.Router, h *handlers.Handler) {
	router.Get("/", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(200) })
	router.Post("/reset-password", h.SendResetPasswordEmailHandler)
	router.Post("/new-password", h.SetNewPasswordHandler)
}

func loadAuthRoutes(router chi.Router, h *handlers.Handler) {
	router.Group(func(router chi.Router) {
		router.Use(custom_middleware.AuthJwt)
		router.Get("/me", h.GetUserHandler)
	})

	router.Group(func(router chi.Router) {
		router.Post("/register", h.CreateUserHandler)
		router.Post("/login", h.LoginUserHandler)
		router.Post("/logout", h.LogoutUserHandler)
	})
}

// Attaches all user releated handlers to chi router
func loadUserRoutes(router chi.Router, h *handlers.Handler) {
	router.Group(func(router chi.Router) {
		router.Use(custom_middleware.AuthJwt)
		router.Get("/contact-info", h.GetUserContactInfo)
		router.Put("/", h.UpdateUserHandler)
		router.Delete("/", h.DeleteUserHandler)
	})
}

// Attaches all client related handlers to chi router
func loadClientRoutes(router chi.Router, h *handlers.Handler) {
	router.Group(func(router chi.Router) {
		router.Use(custom_middleware.AuthJwt)
		router.Post("/", h.CreateClientHandler)
		router.Get("/", h.GetClientHandler)
		router.Put("/", h.UpdateClientHandler)
		router.Put("/archive", h.BatchArchiveClientsHandler)
		router.Delete("/", h.DeleteClientHandler)
	})
}

// Attaches all event related handlers to chi router
func loadEventRoutes(router chi.Router, h *handlers.Handler) {
	router.Group(func(router chi.Router) {
		router.Use(custom_middleware.AuthJwt)
		router.Post("/", h.CreateEventHandler)
		router.Get("/", h.GetEventHandler)
		router.Put("/", h.UpdateEventHandler)
		router.Delete("/", h.DeleteEventHandler)
	})
}

func loadPayoutRoutes(router chi.Router, h *handlers.Handler) {
	router.Group(func(router chi.Router) {
		// TODO add middleware to check ownership of payout
		router.Use(custom_middleware.AuthJwt)
		router.Get("/make", h.MakePayoutHandler)
		router.Post("/", h.SavePayoutHandler)
		router.Get("/", h.GetPayoutsHandler)
		router.Delete("/", h.DeletePayoutHandler)
	})
}

// Attaches all event type related handlers to chi router
func loadEventTypeRoutes(router chi.Router, h *handlers.Handler) {
	router.Group(func(router chi.Router) {
		router.Use(custom_middleware.AuthJwt)
		router.Post("/", h.CreateEventTypeHandler)
		router.Get("/", h.GetEventTypeHandler)
		router.Put("/", h.UpdateEventTypeHandler)
		router.Delete("/", h.DeleteEventTypeHandler)
	})
}

// Attaches all payment type related handlers to router
func loadPaymentTypeRoutes(router chi.Router, h *handlers.Handler) {
	router.Group(func(router chi.Router) {
		router.Use(custom_middleware.AuthJwt)
		router.Post("/", h.CreatePaymentTypeHandler)
		router.Get("/", h.GetPaymentTypesHandler)
		router.Put("/", h.UpdatePaymentTypeHandler)
		router.Delete("/", h.DeletePaymentTypeHandler)
	})
}

// Attaches all global middleware, intended for the main router
func attachMiddleware(router chi.Router) {
	router.Use(middleware.Recoverer)
	router.Use(middleware.Throttle(100))
	router.Use(middleware.Timeout(30 * time.Second))
	// Logging
	router.Use(middleware.Logger)
	// CORS
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173", "http://localhost:4173", "http://localhost:6006", os.Getenv("CLIENT_HOST")},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	}))
}
