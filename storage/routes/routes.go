package handlers

import (
	"encoding/json"
	"log"

	config "github.com/The-Origin-Labs/landate/config"
	handler "github.com/The-Origin-Labs/landate/storage/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func StorageService() {

	// fiber client
	// Initialize a new Fiber application
	// with custom JSON encoding and decoding.
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	// middlewares
	// Enable CORS middleware for cross-origin requests.
	app.Use(cors.New())

	// Add a unique request ID to each incoming request.
	app.Use(requestid.New())

	// Use the built-in logger middleware to log information about each request.
	// For more options, see the Config section
	app.Use(logger.New(logger.Config{
		Format: `${pid} ${locals:requestid} ${status} - ${method} ${path}\n`,
	}))

	// Define a route for the root path to return a simple JSON response.
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"message": "Welcome to Heimdal API Service.",
		})
	})

	// API version: v1
	// Create a new route group for version 1 of the API.
	v1 := app.Group("/v1")
	v1.Get("/", handler.Init)
	v1.Post("/add", handler.AddProperty)
	v1.Get("/all", handler.GetAllProperties)
	v1.Get("/:id", handler.GetProperty)
	v1.Delete("/:id", handler.DeleteProperty)

	// Get the listening port from
	// the environment configuration.
	PORT := config.GetEnvConfig("STORAGE_SERVICE_PORT")

	// Start listening and serving
	// the Fiber application on the specified port.
	log.Fatal(app.Listen(":" + PORT))
}
