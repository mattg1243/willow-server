package application

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/mattg1243/willow-server/handlers"
	custom_middleware "github.com/mattg1243/willow-server/middleware"
)

func loadRoutes(h *handlers.Handler) *chi.Mux {
	// Create new chi router
	router := chi.NewRouter()
	// Global middleware
	router.Use(middleware.Logger)
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
	router.Mount("/event-type", eventTypeRouter)
	// Return the completed router
	return router
}

// Attaches all user releated handlers to chi router
func loadUserRoutes(router chi.Router, h *handlers.Handler) {
// no auth
	router.Group(func(router chi.Router) {
		router.Post("/", h.CreateUserHandler)
		router.Post("/login", h.LoginUserHandler)
	})
// auth
	router.Group(func(router chi.Router) {
		router.Use(custom_middleware.AuthJwt)
		router.Get("/", h.GetUserHandler)
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
		router.Put("/", h.CreateClientHandler)
		router.Delete("/", h.DeleteClientHandler)
	})
}

// Attaches all event related handlers to chi router
func loadEventRoutes(router chi.Router, h *handlers.Handler){
	router.Group(func(router chi.Router) {
		router.Use(custom_middleware.AuthJwt)
		router.Post("/", h.CreateEventHandler)
		router.Get("/", h.GetEventHandler)
		router.Put("/", h.UpdateEventHandler)
		router.Delete("/", h.DeleteEventHandler)
	})
}

// Attaches all event type related handlers to chi router
func loadEventTypeRoutes (router chi.Router, h *handlers.Handler) {

}