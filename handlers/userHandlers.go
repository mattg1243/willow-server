package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/mattg1243/sqlc-fiber/db"
)


func (h *Handler) GetUsersHandler(c *fiber.Ctx) error {
	users, err := h.queries.GetUsers(c.Context())
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.JSON(users)
}

func (h *Handler) GetUserHandler(c *fiber.Ctx) error {
	user, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(200).JSON(user)
}

func (h *Handler) CreateUserHandler(c* fiber.Ctx) error {
	var user db.User
	req := &createUserRequest{}
	
	if err := req.bind(c, &user, h.validator); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(err.Error())
	}

	newUser, err := h.queries.CreateUser(c.Context(), db.CreateUserParams{Username: user.Username, Email: user.Email, Balance: user.Balance})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(http.StatusCreated).JSON(newUser)

}

func (h *Handler) UpdateUserHandler(c *fiber.Ctx) error {
	var user db.User
	userId, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}
	req := &updateUserRequest{}

	if err := req.bind(c, &user, h.validator); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(err.Error())
	}

	updatedUser, err := h.queries.UpdateUser(c.Context(), db.UpdateUserParams{ID: int32(userId), Username: user.Username, Email: user.Email})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.JSON(updatedUser)
}

func (h *Handler) DeleteUserHandler(c *fiber.Ctx) error {
	albumId, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	err = h.queries.DeleteAlbum(c.Context(), int32(albumId))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(200).JSON("User deleted successfully")
}