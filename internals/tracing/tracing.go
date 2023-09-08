package tracing

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	oteltrace "go.opentelemetry.io/otel/trace"
)

var (
	tp     *sdktrace.TracerProvider
	tracer = otel.Tracer("auth-server")
)

func InitTracer() error {
	exporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	if err != nil {
		return fmt.Errorf("failed to initialize stdouttrace exporter: %w", err)
	}

	bsp := sdktrace.NewBatchSpanProcessor(exporter)
	resource := sdktrace.WithResource(
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("trace-svc"),
		))
	tp = sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithSpanProcessor(bsp),
		resource,
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	))
	return nil
}

func CreateSpan(ctx context.Context, spanName string, attr string) string {
	_, span := tracer.Start(ctx,
		spanName,
		oteltrace.WithAttributes(
			attribute.String("attr", attr),
		))
	defer span.End()
	if attr == "" {
		return "unknown"
	}

	return "otel tester"
}
