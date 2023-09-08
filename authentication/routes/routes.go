package routes

import (
	"context"
	"fmt"
	handler "landate/authentication/handlers"
	config "landate/config"
	"landate/internals/tracing"
	"net/http"
	"time"

	"github.com/gofiber/contrib/otelfiber"
	"github.com/gofiber/fiber/v2"
	adaptor "github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

var tracer = otel.Tracer("landate/authentication/routes")

// sleepy mocks work that your application does.
func sleepy(ctx context.Context) {
	_, span := tracer.Start(ctx, "sleep")
	defer span.End()

	sleepTime := 1 * time.Second
	time.Sleep(sleepTime)
	span.SetAttributes(attribute.Int("sleep.duration", int(sleepTime)))
}

// httpHandler is an HTTP handler function that is going to be instrumented.
func httpHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World! I am instrumented automatically!")
	ctx := r.Context()
	sleepy(ctx)
}

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
	_ = tracing.InitTracer()
	app := fiber.New()

	app.Use(logger.New())
	app.Use(otelfiber.Middleware())

	handler := http.HandlerFunc(httpHandler)
	wrappedHandler := otelhttp.NewHandler(handler, "otel-instrumented")
	app.Get("/api/traces", adaptor.HTTPHandler(wrappedHandler))

	AuthRoutes(app)

	AUTH_PORT := config.GetEnvConfig("AUTH_SERVICE_PORT")
	err := app.Listen(":" + AUTH_PORT)
	return err
}
