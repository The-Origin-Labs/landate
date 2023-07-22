package middlewares

import (
	"landate/config"

	"github.com/gofiber/fiber/v2"
)

const (
	authHeader = "AuthKey"
)

func ValidateAPIKey(c *fiber.Ctx) error {
	// Gets AuthKey Header
	authkey := c.Get(authHeader)
	API_ACCESS_KEY := config.GetEnvConfig("API_ACCESS_KEY")

	if authkey != API_ACCESS_KEY {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Welcome to Landate API gateway.",
			"error":   "Need API key to Access. (Invalid API key)",
		})
	}

	return c.Next()
}
