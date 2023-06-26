package handlers

import (
	"log"
	"strings"

	v1Handlers "landate/api-gateway/handlers/versions"
	"landate/api-gateway/middlewares"
	route "landate/api-gateway/routes"
	config "landate/config"

	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

func Gateway() {

	router := fiber.New()

	// API Authorization Middleware
	router.Use(middlewares.ValidateAPIKey)

	router.Use(logger.New(logger.Config{
		Format: `${pid} ${locals:requestid} ${status} - ${method} ${path} /n`,
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

	// connecting storage micro-service
	router.All("/storage/*", func(c *fiber.Ctx) error {
		path := strings.Split(c.Path(), "/storage")[1]
		storageURL := "http://localhost:" + config.GetEnvConfig("STORAGE_SERVICE_PORT")
		return proxy.Do(c, storageURL+path)
	})

	// connecting document microservice
	router.All("/document/*", func(c *fiber.Ctx) error {
		path := strings.Split(c.Path(), "/document")[1]
		docURL := "http://localhost:" + config.GetEnvConfig("DOCUMENT_SERVICE_PORT")
		return proxy.Do(c, docURL+path)
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
