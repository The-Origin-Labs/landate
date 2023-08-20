package routes

import (
	handler "landate/authentication/handlers"
	"landate/config"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(router *fiber.App) {
	router.Get("/", handler.Init)
	router.Post("/auth/user", handler.CreateUser)
	router.Get("/auth/user/:id", handler.RetrieveUser)
	router.Delete("/auth/user/:id", handler.RemoveUser)
	router.Get("/auth/users", handler.GetAllUsers)
	router.Patch("/auth/user/:id", handler.UpdateUser)
	router.Get("/auth/user/:walletAddress", handler.GetUserByWalledId)

}

func Init() error {

	app := fiber.New()

	AuthRoutes(app)

	AUTH_PORT := config.GetEnvConfig("AUTH_SERVICE_PORT")
	err := app.Listen(":" + AUTH_PORT)
	return err
}
