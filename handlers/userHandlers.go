package handlers

import (
	"github.com/gofiber/fiber/v2"
)

// mock db table
type User struct {
	Id uint `json:"id"`
	Username string `json:"username"`
	Email string`json:"email"`
}

func NewUser(username string, email string) User {
	return User{1, username, email}
}

func (h *Handler) GetUserHandler(c *fiber.Ctx) error {
	user := NewUser("mattg1243", "mattgallucci97@gmail.com")
	return c.Status(200).JSON(user)
}