package handlers

import (
	"encoding/json"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/mattg1243/sqlc-fiber/db"
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
	ContactInfo struct {
		Phone string `json:"phone"`
		City string `json:"city"`
		State string `json:"state"`
		Street string `json:"street"`
		Zip string `json:"zip"`
	} `json:"contactInfo"`
}

func (r *createUserRequest) bind(c *fiber.Ctx, u *db.User, cI *db.UserContactInfo, v *Validator) error {
	// validate
	if err := c.BodyParser(r); err != nil {
		return err
	}

	if err := v.Validate(r); err != nil {
		return err
	}
	// bind user
	u.Fname = r.User.Fname
	u.Lname = r.User.Lname
	u.Email = r.User.Email
	// hash password
	h, err := u.HashPassword(r.User.Password)
	if err != nil {
		return err
	}
	u.Hash = h
	// save contact info
	cI.Phone = pgtype.Text{String: r.ContactInfo.Phone, Valid: true}
	cI.City = pgtype.Text{String: r.ContactInfo.City, Valid: true}
	cI.State = pgtype.Text{String: r.ContactInfo.State, Valid: true}
	cI.Street = pgtype.Text{String: r.ContactInfo.Street, Valid: true}
	cI.Zip = pgtype.Text{String: r.ContactInfo.Zip, Valid: true}

	return nil
}

// client requests
type createClientRequest struct {
	Client struct {
		Fname string `json:"fname" validate:"required"`
		Lname string `json:"lname"`
		Email string `json:"email"`
		Rate  int16  `json:"rate" validate:"required"`
		Phone string `json:"phone"`
		Balancenotifythreshold int16 `json:"balanceNotifyThreshold"`
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
	cl.Phone = pgtype.Text{String: r.Client.Phone, Valid: true}
	cl.Balancenotifythreshold = int32(r.Client.Balancenotifythreshold)

	return nil
}

type updateUserRequest struct {
	User struct {
		Fname    string `json:"fname" validate:"required"`
		Lname    string `json:"lname" validate:"required"`
		License	 string `json:"license" validate:"required"`
		Nameforheader string `json:"nameForHeader" validate:"required"`
	} `json:"user"`
	ContactInfo struct {
		Phone string `json:"phone" validate:"required"`
		City string `json:"city" validate:"required"`
		State string `json:"state" validate:"required"`
		Street string `json:"street" validate:"required"`
		Zip string `json:"zip" validate:"required"`
		PaymentInfo struct {
			Venmo string `json:"venmo"`
			Paypal string `json:"paypal"`
		} `json:"paymentInfo"`
	} `json:"contactInfo"`
}

func (r *updateUserRequest) bind(c *fiber.Ctx, u *db.User, cI *db.UserContactInfo, v *Validator) error {
	// validate
	if err := c.BodyParser(r); err != nil {
		return err
	}

	if err := v.Validate(r); err != nil {
		return err
	}
	// bind user
	u.Fname = r.User.Fname
	u.Lname = r.User.Lname
	u.Nameforheader = r.User.Nameforheader
	u.License = pgtype.Text{String: r.User.License, Valid: true}
	// bind contact info
	cI.Phone = pgtype.Text{String: r.ContactInfo.Phone, Valid: true}
	cI.City = pgtype.Text{String: r.ContactInfo.City, Valid: true}
	cI.State = pgtype.Text{String: r.ContactInfo.State, Valid: true}
	cI.Street = pgtype.Text{String: r.ContactInfo.Street, Valid: true}
	cI.Zip = pgtype.Text{String: r.ContactInfo.Zip, Valid: true}
	paymentInfo, err := json.Marshal(r.ContactInfo.PaymentInfo)
	if err != nil {
		return err
	}
	cI.Paymentinfo = paymentInfo


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
	cl.UpdatedAt = pgtype.Timestamp{Time: updatedAtStr}

	return nil
}
