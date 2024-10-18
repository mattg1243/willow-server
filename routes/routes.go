package routes

// import (
// 	"github.com/go-chi/chi/v5"
// 	"github.com/mattg1243/willow-server/handlers"
// 	"github.com/mattg1243/willow-server/middleware"
// )

// func LoadRoutes(r chi.Router, h *handlers.Handler) {
// 	// root
// 	a.Get("/", h.GetRootHandler)

// 	// user routes
// 	userRoutes := r.Group("/user")

// 	userRoutes.Post("/", h.CreateUserHandler)
// 	userRoutes.Post("/login", h.LoginUserHandler)
// 	userRoutes.Get("/", middleware.AuthJwt, h.GetUserHandler)
// 	userRoutes.Get("/contact-info", middleware.AuthJwt, h.GetUserContactInfo)
// 	userRoutes.Put("/", middleware.AuthJwt, h.UpdateUserHandler)
// 	userRoutes.Delete("/", middleware.AuthJwt, h.DeleteUserHandler)

// 	// client routes
// 	clientRoutes := a.Group("/client")

// 	clientRoutes.Post("/", middleware.AuthJwt, h.CreateClientHandler)
// 	clientRoutes.Get("/", middleware.AuthJwt, h.GetClientHandler)
// 	// TODO: test route, impl middleware hook
// 	clientRoutes.Put("/", middleware.AuthJwt, h.UpdateClientHandler)
// 	clientRoutes.Delete("/", middleware.AuthJwt, h.DeleteClientHandler)

// 	// event routes
// 	eventRoutes := a.Group("/event")
// 	eventRoutes.Post("/", h.CreateEventHandler)

// 	// event type routes
// 	eventTypeRoutes := a.Group("/event-types")
// 	eventTypeRoutes.Post("/", middleware.AuthJwt, h.CreateEventTypeHandler)

// }
