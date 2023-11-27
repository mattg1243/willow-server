package handlers

import "github.com/gofiber/fiber/v2"

func (h *Handler) GetRootHandler(c *fiber.Ctx) error {
	return c.Status(200).JSON("Server online!")
}