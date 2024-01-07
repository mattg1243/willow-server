package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/mattg1243/sqlc-fiber/db"
)

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

	newUser, err := h.queries.CreateUser(c.Context(), db.CreateUserParams{ Hash: user.Hash, Email: user.Email})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	newUser.Hash = ""

	return c.Status(http.StatusCreated).JSON(newUser)

}

func (h *Handler) UpdateUserHandler(c *fiber.Ctx) error {
	var user db.User
	req := &updateUserRequest{}

	if err := req.bind(c, &user, h.validator); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(err.Error())
	}

	updatedUser, err := h.queries.UpdateUser(c.Context(), db.UpdateUserParams{
		Fname: req.User.Fname,
		Lname: req.User.Lname,
		Phone: pgtype.Text{String: req.User.Phone},
		Nameforheader: req.User.NameForHeader,
		Street: pgtype.Text{String: req.User.Street},
		City: pgtype.Text{String: req.User.City},
		Zip: pgtype.Text{String: req.User.Zip},
		State: pgtype.Text{String: req.User.State},
		License: pgtype.Text{String: req.User.License},
		Paymentinfo: []byte(req.User.PyamentInfo.PayPal),
	})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.JSON(updatedUser)
}

// func (h *Handler) DeleteUserHandler(c *fiber.Ctx) error {
// 	userId, err := c.ParamsInt("id")

// 	if err != nil {
// 		return c.Status(http.StatusBadRequest).JSON(err.Error())
// 	}

// 	return c.Status(200).JSON("User deleted successfully")
// }

// func (h *Handler) LoginUserHandler(c *fiber.Ctx) error {
// 	req := loginUserRequest{}

// 	if err := req.bind(c, h.validator); err != nil {
// 		return c.Status(http.StatusBadRequest).JSON(err.Error())
// 	}

// 	user, err := h.queries.GetUserWithHash(c.Context(), req.Username)
// 	if err != nil {
// 		return c.Status(http.StatusInternalServerError).JSON(err.Error())
// 	}

// 	match := user.CheckPassword(req.Password)

// 	if (match) {
// 		payload := utils.JwtPayload{Id: user.ID, Username: user.Username, Email: user.Email}
// 		jwt, err := utils.GenerateJWT(payload)
// 		if err != nil {
// 			log.Fatalf(err.Error())
// 			return c.Status(http.StatusInternalServerError).JSON(err.Error())
// 		}
// 		c.Cookie(&fiber.Cookie{
// 			Name: "access-token",
// 			Expires: time.Now().Add((time.Hour * 72)),
// 			HTTPOnly: false,
// 			Secure: false,
// 			SameSite: "lax",
// 			Value: jwt,
// 		})

// 		return c.SendStatus(200)
// 	} else {
// 		return c.Status(http.StatusUnauthorized).JSON("Invalid login credentials")
// 	}
// }