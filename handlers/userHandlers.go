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
	userIdStr := c.Locals("user")
	userId, err := uuid.Parse(userIdStr.(string))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	user, err := h.queries.GetUser(c.Context(), userId)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(200).JSON(user)
}

func (h *Handler) GetUserContactInfo(c *fiber.Ctx) error {
	userIdStr := c.Locals("user")
	userId, err := uuid.Parse(userIdStr.(string))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	contactInfo, err := h.queries.GetUserContactInfo(c.Context(), userId)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(200).JSON(contactInfo)
}

func (h *Handler) CreateUserHandler(c *fiber.Ctx) error {
	var user db.User
	var contactInfo db.UserContactInfo
	req := &createUserRequest{}

	if err := req.bind(c, &user, &contactInfo, h.validator); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(err.Error())
	}
	// hash the password
	hash, err := user.HashPassword(req.User.Password)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}
	// save user
	newUser, err := h.queries.CreateUser(c.Context(), db.CreateUserParams{ 
		Hash: hash, 
		Email: user.Email,
		Fname: user.Fname,
		Lname: user.Lname,
		Nameforheader: user.Nameforheader,
		ID: uuid.New(),
	})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}
	// save contact info
	_, err = h.queries.CreateUserContactInfo(c.Context(), db.CreateUserContactInfoParams{
		ID: uuid.New(),
		Phone: contactInfo.Phone,
		City: contactInfo.City,
		State: contactInfo.State,
		Street: contactInfo.Street,
		Zip: contactInfo.Zip,
		UserID: newUser.ID,
	})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(http.StatusCreated).JSON(newUser)
}

func (h *Handler) UpdateUserHandler(c *fiber.Ctx) error {
	var user db.User
	var contactInfo db.UserContactInfo
	req := &updateUserRequest{}

	// parse user id from claims
	userIdStr := c.Locals("user")
	userId, err := uuid.Parse(userIdStr.(string))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	if err := req.bind(c, &user, &contactInfo, h.validator); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(err.Error())
	}
	// save user
	updatedUser, err := h.queries.UpdateUser(c.Context(), db.UpdateUserParams{
		Fname: user.Fname,
		Lname: user.Lname,
		Nameforheader: user.Nameforheader,
		License: user.License,
		ID: userId,
	})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}
	// save contact info
	_, err = h.queries.UpdateUserContactInfo(c.Context(), db.UpdateUserContactInfoParams{
		UserID: userId,
		Phone: contactInfo.Phone,
		City: contactInfo.City,
		State: contactInfo.State,
		Street: contactInfo.Street,
		Zip: contactInfo.Zip,
		Paymentinfo: contactInfo.Paymentinfo,
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

	if match {
		payload := utils.JwtPayload{Id: user.ID.String(), Email: user.Email}
		jwt, err := utils.GenerateJWT(payload)
		if err != nil {
			log.Fatalf(err.Error())
			return c.Status(http.StatusInternalServerError).JSON(err.Error())
		}
		c.Cookie(&fiber.Cookie{
			Name:     "willow-access-token",
			Expires:  time.Now().Add((time.Hour * 72)),
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
