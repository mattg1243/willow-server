package handlers

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
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
		Fname                  string `json:"fname" validate:"required"`
		Lname                  string `json:"lname"`
		Email                  string `json:"email"`
		Rate                   int16  `json:"rate" validate:"required"`
		Phone                  string `json:"phone"`
		Balancenotifythreshold int16  `json:"balanceNotifyThreshold"`
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
	Email    string `json:"email" validate:"required,email"`
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

// event requests
type createEventRequest struct {
	Event struct {
		ClientID   uuid.UUID `json:"clientId"`
		Date       string    `json:"date"`
		Duration   float64   `json:"duration"`
		Type       string    `json:"type"`
		Detail     string    `json:"detail"`
		Rate       int32     `json:"rate"`
		Amount     float64   `json:"amount"`
		Newbalance float64   `json:"newbalance"`
	} `json:"event"`
}

func Float64ToPgNumeric(f float64) pgtype.Numeric {
	var n pgtype.Numeric
	if err := n.Scan(fmt.Sprintf("%f", f)); err != nil {
		log.Error("error scanning float64 to pg numeric: ", err)
	}
	return n
}

func (r *createEventRequest) bind(c *fiber.Ctx, e *db.Event, v *Validator) error {
	log.Info("binding req for: event")
	// validate
	if err := c.BodyParser(r); err != nil {
		return err
	}

	if err := v.Validate(r); err != nil {
		return err
	}

	timeLayout := "2006-01-02 15:04:05"
	timeStr, err := time.Parse(timeLayout, r.Event.Date)
	if err != nil {
		log.Error("error parsing time: ", err)
		return err
	}

	e.Date = pgtype.Timestamp{Time: timeStr, Valid: true}
	e.Duration = Float64ToPgNumeric(r.Event.Duration)
	e.Type = pgtype.Text{String: r.Event.Type}
	e.Detail = pgtype.Text{String: r.Event.Detail}
	e.Rate = r.Event.Rate
	e.Amount = Float64ToPgNumeric(r.Event.Amount)
	e.ClientID = r.Event.ClientID
	log.Info("ClientID: ", e.ClientID)
	e.Newbalance = Float64ToPgNumeric(r.Event.Newbalance)

	return nil
}
