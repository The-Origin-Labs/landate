package routes

import (
	"encoding/json"

	"github.com/The-Origin-Labs/landate/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func documentRoutes(route *fiber.App) {

	route.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Welcome to Docuement Microserice ⚡",
		})
	})
}

func Init() error {

	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	// Cross-Origin Resource Sharing
	app.Use(cors.New())

	// Assigning Routes
	documentRoutes(app)

	PORT := config.GetEnvConfig("DOCUMENT_SERVICE_PORT")
	err := app.Listen(":" + PORT)
	return err
}
