package routes

import (
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {

	api := app.Group("/api")

	api.Get("/mv1", func(c *fiber.Ctx) error {
		return c.SendString("Api running...")
	})

}
