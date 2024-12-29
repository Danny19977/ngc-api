package middlewares

import (
	"github.com/Danny19977/ngc-api/utils"
	"github.com/gofiber/fiber/v2"
)

func IsAuthenticated(c *fiber.Ctx) error {

	cookie := c.Cookies("token")

	if _, err := utils.VerifyJwt(cookie); err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	return c.Next()
}