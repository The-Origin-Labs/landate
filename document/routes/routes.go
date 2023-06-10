package routes

import (
	"encoding/json"

	"github.com/The-Origin-Labs/landate/config"
	handler "github.com/The-Origin-Labs/landate/document/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func documentRoutes(route *fiber.App) {
	route.Get("/", handler.Init)
	api := route.Group("/api/user")
	api.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"message": "api/user",
		})
	})
	api.Post("/new", handler.AddUserDocs)
	api.Get("/:id", handler.GetUserDocsById)
	api.Patch("/update/:id", handler.UpdateUserDocs)
	api.Delete("/:id", handler.RemoveUserDocs)
	api.Post("/txid", handler.GetUserDocsByTxId)
	api.Post("/walletaddr", handler.GetUserDocsByWalletAddress)
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
