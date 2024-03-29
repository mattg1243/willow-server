package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mattg1243/sqlc-fiber/handlers"
	"github.com/mattg1243/sqlc-fiber/middleware"
)

func LoadRoutes(a *fiber.App, h *handlers.Handler) {
	// root
	a.Get("/", h.GetRootHandler)

	// user routes
	userRoutes := a.Group("/user")

	userRoutes.Post("/", h.CreateUserHandler)
	userRoutes.Post("/login", h.LoginUserHandler)
	userRoutes.Get("/", middleware.AuthJwt, h.GetUserHandler)
	userRoutes.Get("/contact-info", middleware.AuthJwt, h.GetUserContactInfo)
	userRoutes.Put("/", middleware.AuthJwt, h.UpdateUserHandler)
	userRoutes.Delete("/", middleware.AuthJwt, h.DeleteUserHandler)

	// userRoutes.Post("/login", h.LoginUserHandler)

	// client routes
	clientRoutes := a.Group("/client")

	clientRoutes.Post("/", middleware.AuthJwt, h.CreateClientHandler)
	clientRoutes.Get("/", middleware.AuthJwt, h.GetClientHandler)
	// TODO: test route, impl middleware hook
	clientRoutes.Put("/", middleware.AuthJwt, h.UpdateClientHandler)
	clientRoutes.Delete("/", middleware.AuthJwt, h.DeleteClientHandler)

	// event routes
	eventRoutes := a.Group("/event")
	eventRoutes.Post("/", middleware.AuthJwt, h.CreateEventHandler)
	eventRoutes.Get("/", middleware.AuthJwt, h.GetEventHandler)
	eventRoutes.Put("/", middleware.AuthJwt, h.UpdateEventHandler)
	eventRoutes.Delete("/", middleware.AuthJwt, h.DeleteEventHandler)

	eventsRoutes := a.Group("/events")
	eventsRoutes.Get("/", middleware.AuthJwt, h.GetEventsByClientHandler)

	// artist routes
	// artistRoutes := a.Group("/artists")

	// artistRoutes.Post("/", h.CreateArtistHandler)
	// artistRoutes.Get("/", h.GetArtistsHandler)
	// artistRoutes.Get("/:id", h.GetArtistHandler)
	// artistRoutes.Put("/:id", h.UpdateArtistHandler)
	// artistRoutes.Delete("/:id", h.DeleteArtistHandler)

	// purchase routes
	// purchaseRoutes := a.Group("/purchases")
	// purchaseRoutes.Use(keyauth.New(keyauth.Config{
	// 	KeyLookup: "cookie:access-token",
	// 	Validator: middleware.AuthJwt,
	// }))
	// purchaseRoutes.Post("/", h.CreatePurchaseHandler)
	// purchaseRoutes.Get("/", h.GetPurchasesHandler)
}
