package middleware

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/mattg1243/sqlc-fiber/utils"
)


func AuthJwt (c *fiber.Ctx) error {
	
	token := c.Cookies("willow-access-token")

	claims, err := utils.ValidateJWT(token)
	if err != nil {
		return errors.New("Invalid credentials")
	}
	
	c.Locals("user", claims.Id)
	c.Locals("email", claims.Email)
	return c.Next()
}