package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/keyauth"
	"github.com/mattg1243/sqlc-fiber/handlers"
	"github.com/mattg1243/sqlc-fiber/middleware"
)

func LoadRoutes(a *fiber.App, h *handlers.Handler) {
	// root routes
	a.Get("/", h.GetRootHandler)
	// album routes
	albumRoutes := a.Group("/albums")

	albumRoutes.Post("/", h.CreateAlbumHandler)
	albumRoutes.Get("/", h.GetAlbumsHandler)
	albumRoutes.Get("/:id", h.GetAlbumHandler)
	albumRoutes.Put("/:id", h.UpdateAlbumHandlder)
	albumRoutes.Delete("/:id", h.DeleteAlbumHandler)
	
	// user routes
	userRoutes := a.Group("/users")

	userRoutes.Post("/", h.CreateUserHandler)
	userRoutes.Get("/", h.GetUsersHandler)
	userRoutes.Get("/:id", h.GetUserHandler)
	userRoutes.Put("/:id", h.UpdateUserHandler)
	userRoutes.Delete("/:id", h.DeleteUserHandler)

	userRoutes.Post("/login", h.LoginUserHandler)

	// artist routes
	artistRoutes := a.Group("/artists")

	artistRoutes.Post("/", h.CreateArtistHandler)
	artistRoutes.Get("/", h.GetArtistsHandler)
	artistRoutes.Get("/:id", h.GetArtistHandler)
	artistRoutes.Put("/:id", h.UpdateArtistHandler)
	artistRoutes.Delete("/:id", h.DeleteArtistHandler)

	// purchase routes
	purchaseRoutes := a.Group("/purchases")
	purchaseRoutes.Use(keyauth.New(keyauth.Config{
		KeyLookup: "cookie:access-token",
		Validator: middleware.AuthJwt,
	}))
	purchaseRoutes.Post("/", h.CreatePurchaseHandler)
	purchaseRoutes.Get("/", h.GetPurchasesHandler)
}