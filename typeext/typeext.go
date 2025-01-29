package typeext

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type PaymentInfo map[string]interface{}


// Scan implements the sql.Scanner interface to convert []byte to JSON
func (p *PaymentInfo) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		if err := json.Unmarshal(v, p); err != nil {
			return fmt.Errorf("PaymentInfo Scan error: %w", err)
		}
	case string: // Handle cases where PostgreSQL sends it as a string
		if err := json.Unmarshal([]byte(v), p); err != nil {
			return fmt.Errorf("PaymentInfo Scan error (string case): %w", err)
		}
	default:
		return fmt.Errorf("PaymentInfo Scan: unexpected type %T", value)
	}
	return nil
}

// Value implements the driver.Valuer interface to store JSON
func (p PaymentInfo) Value() (driver.Value, error) {
	return json.Marshal(p)
}