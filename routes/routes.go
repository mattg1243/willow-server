package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mattg1243/sqlc-fiber/handlers"
)

func LoadRoutes(a *fiber.App, h *handlers.Handler) {
	// root
	a.Get("/", h.GetRootHandler)
	
	// user routes
	userRoutes := a.Group("/users")

	userRoutes.Post("/", h.CreateUserHandler)
	userRoutes.Get("/:id", h.GetUserHandler)
	userRoutes.Put("/:id", h.UpdateUserHandler)
	userRoutes.Delete("/:id", h.DeleteUserHandler)

	userRoutes.Post("/login", h.LoginUserHandler)

	// client routes
	clientRoutes := a.Group("/clients")

	clientRoutes := a.Post("/", h.)

	// event routes

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