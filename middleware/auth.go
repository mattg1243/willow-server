package middleware

import (
	"errors"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/mattg1243/sqlc-fiber/utils"
)



func AuthJwt (c *fiber.Ctx) error {
	key := os.Getenv("JWT_SECRET")
	claims, err := utils.ValidateJWT(key)
	if err != nil {
		return errors.New("Invalid credentials")
	}

	c.Locals("jwtClaims", claims)
	return c.Next()
}