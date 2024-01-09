package db

import (
	"encoding/json"
)

func (u *UserContactInfo) PaymentInfoToString() []byte {
	str, err := json.Marshal(u.Paymentinfo)

	if err != nil {
		return nil
	}

	return str
}