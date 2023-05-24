package handlers

import (
	"log"

	v1Handlers "github.com/The-Origin-Labs/landate/api-gateway/handlers/versions"
	route "github.com/The-Origin-Labs/landate/api-gateway/routes"
	config "github.com/The-Origin-Labs/landate/config"

	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Gatway() {

	router := fiber.New()

	router.Use(logger.New(logger.Config{
		Format: `${pid} ${locals:requestid} ${status} - ${method} ${path}\n`,
	}))

	router.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"uri":     ctx.Request().URI().String(),
			"path":    ctx.Path(),
			"message": "Welcome To API Gateway",
		})
	})

	router.Get("/env", func(c *fiber.Ctx) error {

		return c.JSON(fiber.Map{
			"uri":  c.Request().URI().String(),
			"port": "env",
		})
	})

	v1 := router.Group("/v1")
	v1.Get("/", v1Handlers.BaseRoute)
	v1.Get("/g", func(c *fiber.Ctx) error {
		return c.SendString("g")
	})

	route.Routes(router)

	// Middlewares
	router.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	// Prometheus Client middleware
	prometheus := fiberprometheus.New("api-gateway-traces")
	prometheus.RegisterAt(router, "/metrics")
	router.Use(prometheus.Middleware)
	PORT := config.GetEnvConfig("API_GATEWAY_PORT")

	log.Fatal(router.Listen(":" + PORT))
}