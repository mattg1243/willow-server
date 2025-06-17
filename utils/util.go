package utils

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func GenerateResetPasswordToken() (string, error) {
  bytes := make([]byte, 32)

  if _, err := rand.Read(bytes); err != nil {
    return "", fmt.Errorf("failed to generate secure random token: %w", err)
  }

  token := base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(bytes)
  return token, nil
}