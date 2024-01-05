package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mattg1243/sqlc-fiber/db"
)

// user requests
type createUserRequest struct {
	User struct {
		Password string `json:"password" validate:"required"`
		Email string `json:"email" validate:"required,email"`
		Fname string `json:"fname" validate:"required"`
		Lname string `json:"lname" validate:"required"`
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

	u.Fname = r.User.Fname
	u.Lname = r.User.Lname
	u.Email = r.User.Email
	// hash password
	h, err := u.HashPassword(r.User.Password)
	if err != nil {
		return err
	}
	u.Hash = h

	return nil
}

// client requests
type createClientRequest struct {
	Client struct {
		Fname string `json:"fname" validate:"required"`
		Lname string `json:"lname"`
		Email string `json:"email"`
		Rate int16 `json:"rate" validate:"required"`
	} `json:"client"`
}

func (r *createClientRequest) bind(c *fiber.Ctx, cl *db.Client, v *Validator) error {
	// validate
	if err := c.BodyParser(r); err != nil {
		return err
	}

	if err := v.Validate(r); err != nil {
		return err
	}

	cl.Fname = r.Client.Fname
	cl.Lname = r.Client.Lname
	cl.Email = r.Client.Email
	cl.Rate = int32(r.Client.Rate)

	return nil
}

// type updateUserRequest struct {
// 	User struct {
// 		Username string `json:"username" validate:"required"`
// 		Email string `json:"email" validate:"required,email"`
// 		Balance int32 `json:"balance" validate:"required,gte=0"`
// 	}
// }

// func (r *updateUserRequest) bind(c *fiber.Ctx, u *db.User, v *Validator) error {
// 	// validate
// 	if err := c.BodyParser(r); err != nil {
// 		return err
// 	}

// 	if err := v.Validate(r); err != nil {
// 		return err
// 	}

// 	u.Username = r.User.Username
// 	u.Email = r.User.Email
// 	u.Balance = r.User.Balance

// 	return nil
// }

// type loginUserRequest struct {
// 	Username string `json:"username" validate:"required"`
// 	Password string `json:"password" validate:"required"`
// }

// func (r *loginUserRequest) bind(c *fiber.Ctx, v *Validator) error {
// 	if err :=c.BodyParser(r); err != nil {
// 		return err
// 	}

// 	if err := v.Validate(r); err != nil {
// 		return err
// 	}

// 	return nil
// }

// TODO create request structs and bind funcs for the routes corresponding to the other models

