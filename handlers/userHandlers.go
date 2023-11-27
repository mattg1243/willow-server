package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/mattg1243/sqlc-fiber/db"
)


func (h *Handler) GetUsersHandler(c *fiber.Ctx) error {
	users, err := h.Queries.GetUsers(c.Context())
	if err != nil {
		log.Fatalf("An error occured:\n%s", err.Error())
		return c.SendStatus(500)
	}

	return c.JSON(users)
}

func (h *Handler) GetUserHandler(c *fiber.Ctx) error {
	user, err := c.ParamsInt("id")
	if err != nil {
		log.Fatalf("An error occured:\n%s", err.Error())
		return c.SendStatus(500)
	}

	return c.Status(200).JSON(user)
}

func (h *Handler) CreateUserHandler(c* fiber.Ctx) error {
 var user db.User

	if err := c.BodyParser(&user); err != nil {
		log.Fatalf("An error occured:\n%s", err.Error())
		return c.Status(500).JSON(err.Error())
	}

	newUser, err := h.Queries.CreateUser(c.Context(), db.CreateUserParams{Username: user.Username, Email: user.Email})
	if err != nil {
		log.Fatalf("An error occured:\n%s", err.Error())
		return c.Status(500).JSON(err.Error())
	}

	return c.Status(200).JSON(newUser)

}

func (h *Handler) UpdateUserHandler(c *fiber.Ctx) error {
	var user db.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	updatedUser, err := h.Queries.UpdateUser(c.Context(), db.UpdateUserParams{ID: user.ID, Username: user.Username, Email: user.Email})
	if err != nil {
		log.Fatalf("An error occured:\n%s", err.Error())
		return c.SendStatus(500)
	}

	return c.JSON(updatedUser)
}

func (h *Handler) DeleteUserHandler(c *fiber.Ctx) error {
	albumId, err := c.ParamsInt("id")

	if err != nil {
		log.Fatalf(err.Error())
		c.Status(400).JSON(err.Error())
	}

	err = h.Queries.DeleteAlbum(c.Context(), int32(albumId))
	if err != nil {
		log.Fatalf(err.Error())
		c.Status(400).JSON(err.Error())
	}

	return c.Status(200).JSON("User deleted successfully")
}