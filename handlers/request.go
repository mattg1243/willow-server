package handlers

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/mattg1243/sqlc-fiber/db"
	"time"
)

type PaymentInfo struct {
	Venmo  string `json:"venmo"`
	PayPal string `json:"paypal"`
}

// user requests
type createUserRequest struct {
	User struct {
		Password string `json:"password" validate:"required"`
		Email    string `json:"email" validate:"required,email"`
		Fname    string `json:"fname" validate:"required"`
		Lname    string `json:"lname" validate:"required"`
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
		Rate  int16  `json:"rate" validate:"required"`
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
	cl.Lname = pgtype.Text{String: r.Client.Lname, Valid: true}
	cl.Email = pgtype.Text{String: r.Client.Email, Valid: true}
	cl.Rate = int32(r.Client.Rate)

	return nil
}

type updateUserRequest struct {
	User struct {
		Fname         string      `json:"fname"`
		Lname         string      `json:"lname"`
		Phone         string      `json:"phone"`
		NameForHeader string      `json:"nameForHeader"`
		Street        string      `json:"street"`
		City          string      `json:"city"`
		Zip           string      `json:"zip"`
		State         string      `json:"state"`
		License       string      `json:"license"`
		PyamentInfo   PaymentInfo `json:"paymentInfo"`
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

	// marshall payment info
	piBytes, err := json.Marshal(r.User.PyamentInfo)

	if err != nil {
		return err
	}

	u.Fname = r.User.Fname
	u.Lname = r.User.Lname
	u.Phone = pgtype.Text{String: r.User.Phone, Valid: true}
	u.Nameforheader = r.User.NameForHeader
	u.Street = pgtype.Text{String: r.User.Street, Valid: true}
	u.City = pgtype.Text{String: r.User.City, Valid: true}
	u.Zip = pgtype.Text{String: r.User.Zip, Valid: true}
	u.State = pgtype.Text{String: r.User.State, Valid: true}
	u.License = pgtype.Text{String: r.User.License, Valid: true}
	u.Paymentinfo = piBytes
	// need to improve typing on this

	// u.Paymentinfo = r.User.Paymentinfo

	return nil
}

type loginUserRequest struct {
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (r *loginUserRequest) bind(c *fiber.Ctx, v *Validator) error {
	if err := c.BodyParser(r); err != nil {
		return err
	}

	if err := v.Validate(r); err != nil {
		return err
	}

	return nil
}

type getClientsRequest struct{}

func (r *getClientsRequest) bind(c *fiber.Ctx, v *Validator) error {
	return nil
}

type updateClientRequest struct {
	Client struct {
		Fname                  string `json:"fname"`
		Lname                  string `json:"lname"`
		Email                  string `json:"email"`
		Balance                int32  `json:"balance"`
		Balancenotifythreshold int32  `json:"balancenotifythreshold"`
		Rate                   int32  `json:"rate"`
		Isarchived             bool   `json:"isArchived"`
		CreatedAt              string `json:"createdAt"`
		UpdateAt               string `json:"updatedAt"`
	}
}

func (r *updateClientRequest) bind(c *fiber.Ctx, cl *db.Client, v *Validator) error {
	if err := c.BodyParser(r); err != nil {
		return err
	}

	if err := v.Validate(r); err != nil {
		return err
	}

	cl.Fname = r.Client.Fname
	cl.Lname = pgtype.Text{String: r.Client.Lname}
	cl.Email = pgtype.Text{String: r.Client.Email}
	cl.Balance = r.Client.Balance
	cl.Balancenotifythreshold = r.Client.Balancenotifythreshold
	cl.Rate = r.Client.Rate
	cl.Isarchived = pgtype.Bool{Bool: r.Client.Isarchived}

	pgTimestampLayout := "2006-01-02 15:04:05.999999-07"

	createdAtStr, err := time.Parse(pgTimestampLayout, r.Client.CreatedAt)
	if err != nil {
		return err
	}
	updatedAtStr, err := time.Parse(pgTimestampLayout, r.Client.UpdateAt)
	if err != nil {
		return err
	}

	cl.CreatedAt = pgtype.Timestamp{Time: createdAtStr}
	cl.UpdateAt = pgtype.Timestamp{Time: updatedAtStr}

	return nil
}
