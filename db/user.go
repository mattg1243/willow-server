package db

import (
	"encoding/json"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type PaymentInfo struct {
	Venmo string `json:"venmo"`
	PayPal string `json:"paypal"`
}

func (u *User) HashPassword(password string) (string, error) {
	if len(password) < 12 {
		return "", errors.New("Password must be more than 12 characters")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Hash), []byte(password))
	return err == nil
}

func (u *User) PaymentInfoToString() []byte {
	str, err := json.Marshal(u.Paymentinfo)

	if err != nil {
		return nil
	}

	return str
}