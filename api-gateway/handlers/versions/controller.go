package versions

import "github.com/gofiber/fiber/v2"

func BaseRoute(c *fiber.Ctx) error {
	return c.SendString("Base router for api v1")
}
