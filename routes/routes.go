package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mattg1243/sqlc-fiber/handlers"
)

func LoadRoutes(a *fiber.App, h *handlers.Handler) {
	// root routes
	a.Get("/", h.GetRootHandler)
	// album routes
	albumRoutes := a.Group("/albums")

	albumRoutes.Get("/", h.GetAlbumsHandler)
	albumRoutes.Post("/", h.CreateAlbumHandler)
	
	// user routes
	userRoutes := a.Group("/user")

	userRoutes.Get("/", h.GetUserHandler)
}