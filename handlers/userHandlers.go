package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/mattg1243/sqlc-fiber/db"
	"github.com/mattg1243/sqlc-fiber/utils"
)

func (h *Handler) GetUserHandler(c *fiber.Ctx) error {
	user, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(200).JSON(user)
}

func (h *Handler) CreateUserHandler(c *fiber.Ctx) error {
	var user db.User
	req := &createUserRequest{}

	if err := req.bind(c, &user, h.validator); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(err.Error())
	}
	// hash the password
	hash, err := user.HashPassword(req.User.Password)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	newUser, err := h.queries.CreateUser(c.Context(), db.CreateUserParams{ 
		Hash: hash, 
		Email: user.Email,
		Fname: user.Fname,
		Lname: user.Lname,
		Nameforheader: user.Nameforheader,
		Phone: user.Phone,
		Street: user.Street,
		City: user.City,
		State: user.State,
		Zip: user.Zip,
		ID: uuid.New(),
	})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}
	// dont send the hash to the client
	newUser.Hash = ""

	return c.Status(http.StatusCreated).JSON(newUser)

}

func (h *Handler) UpdateUserHandler(c *fiber.Ctx) error {
	var user db.User
	req := &updateUserRequest{}

	// parse user id from claims
	userIdStr := c.Locals("user")
	userId, err := uuid.Parse(userIdStr.(string))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	if err := req.bind(c, &user, h.validator); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(err.Error())
	}

	updatedUser, err := h.queries.UpdateUser(c.Context(), db.UpdateUserParams{
		Fname: user.Fname,
		Lname: user.Lname,
		Phone: user.Phone,
		Nameforheader: user.Nameforheader,
		Street: user.Street,
		City: user.City,
		Zip: user.Zip,
		State: user.State,
		License: user.License,
		Paymentinfo: user.Paymentinfo,
		ID: userId,
	})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.JSON(updatedUser)
}

func (h *Handler) DeleteUserHandler(c *fiber.Ctx) error {
	userIdStr := c.Locals("user").(string)
	userId, err := uuid.Parse(userIdStr)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	queryErr := h.queries.DeleteUser(c.Context(), userId)
	if queryErr != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(200).JSON("User deleted successfully")
}

func (h *Handler) LoginUserHandler(c *fiber.Ctx) error {
	
	req := loginUserRequest{}

	if err := req.bind(c, h.validator); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	user, err := h.queries.GetUserByEmail(c.Context(), req.Email)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	match := user.CheckPassword(req.Password)

	if (match) {
		payload := utils.JwtPayload{Id: user.ID.String(), Email: user.Email}
		jwt, err := utils.GenerateJWT(payload)
		if err != nil {
			log.Fatalf(err.Error())
			return c.Status(http.StatusInternalServerError).JSON(err.Error())
		}
		c.Cookie(&fiber.Cookie{
			Name: "willow-access-token",
			Expires: time.Now().Add((time.Hour * 72)),
			HTTPOnly: false,
			Secure:   false,
			SameSite: "lax",
			Value:    jwt,
		})

		return c.SendStatus(200)
	} else {
		return c.Status(http.StatusUnauthorized).JSON("Invalid login credentials")
	}
}
