package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mattg1243/sqlc-fiber/db"
)

// user requests
type createUserRequest struct {
	User struct {
		Username string `json:"username" validate:"required"`
		Email string `json:"email" validate:"required,email"`
		Balance int32 `json:"balance" validate:"required gte=0"`
	} `json:"user"`
}

func (r *createUserRequest) bind(c *fiber.Ctx, u *db.User, v *Validator) error {
	// validate
	if err := c.BodyParser(r); err != nil {
		return err
	}

	if err := v.Validate(r); err != nil {
		return err
	}

	u.Username = r.User.Username
	u.Email = r.User.Email
	u.Balance = r.User.Balance

	return nil
}

type updateUserRequest struct {
	User struct {
		Username string `json:"username" validate:"required"`
		Email string `json:"email" validate:"required,email"`
		Balance int32 `json:"balance" validate:"required,gte=0"`
	}
}

func (r *updateUserRequest) bind(c *fiber.Ctx, u *db.User, v *Validator) error {
	// validate
	if err := c.BodyParser(r); err != nil {
		return err
	}

	if err := v.Validate(r); err != nil {
		return err
	}

	u.Username = r.User.Username
	u.Email = r.User.Email
	u.Balance = r.User.Balance

	return nil
}

// TODO create request structs and bind funcs for the routes corresponding to the other models

// artist requests

// album requests

// purhcase requests