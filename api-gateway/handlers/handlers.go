package handlers

import (
	"log"
	"strings"
	"time"

	v1Handlers "landate/api-gateway/handlers/versions"
	"landate/api-gateway/middlewares"
	route "landate/api-gateway/routes"
	config "landate/config"

	// "github.com/ansrivas/fiberprometheus/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/proxy"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "http_request_duration_seconds_2",
		Help: "The total number of processed events",
	})
)

func recordMetrics() {
	go func() {
		for {
			opsProcessed.Inc()
			time.Sleep(2 * time.Second)
		}
	}()
}

func Gateway() {

	router := fiber.New()

	// API Authorization Middleware
	router.Use(middlewares.ValidateAPIKey)

	// Previous Feature
	// promefiber := fiberprometheus.New("http_request_duration_seconds")
	// promefiber.RegisterAt(router, "/metrics")
	// router.Use(promefiber.Middleware)

	// Prometheus Configuration
	prouter := fiber.New()
	requestDuration := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Histogram of the request duration.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "route", "status"},
	)
	prometheus.MustRegister(requestDuration)
	prouter.Use(func(c *fiber.Ctx) error {
		start := time.Now()
		err := c.Next()
		duration := time.Since(start).Seconds()
		requestDuration.WithLabelValues(c.Method(), c.Path(), string(c.Response().StatusCode())).Observe(duration)
		return err
	})
	recordMetrics()
	prouter.Get("/metrics", adaptor.HTTPHandler(promhttp.Handler()))
	go func() {
		log.Fatal(prouter.Listen(":2222"))
	}()
	// ========== END of PROMETHEUS Middleware ==========

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

	PORT := config.GetEnvConfig("API_GATEWAY_PORT")

	log.Fatal(router.Listen(":" + PORT))
}
