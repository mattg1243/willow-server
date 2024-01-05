package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/mattg1243/sqlc-fiber/utils"
)

func AuthJwt (c *fiber.Ctx, key string) (bool, error) {
	fmt.Print(c.Cookies("access-token"))
	claims, err := utils.ValidateJWT(key)
	if err != nil {
		return false, c.Status(401).JSON(err.Error())
	}
	c.Locals("user", claims)
	return true, nil
}