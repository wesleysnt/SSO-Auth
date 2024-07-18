package otel

import (
	"context"
	"log"
	"os"
	"strconv"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	sdkTrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"go.opentelemetry.io/otel/trace"
)

var Tracer trace.Tracer
var Tp *sdkTrace.TracerProvider

func Init(ctx context.Context, endPoint string) {
	// init otel tracer
	trace := false
	trace, _ = strconv.ParseBool(os.Getenv("JAEGER_TRACE"))

	if trace {
		exp, err := newExporter(ctx)
		if err != nil {
			log.Fatalf("failed to initialize exporter: %v", err)
		}
		// Create a new tracer provider with a batch span processor and the given exporter.
		Tp = newTraceProvider(exp)

		// Handle shutdown properly so nothing leaks.
		// defer func() { _ = tp.Shutdown(ctx) }()

		otel.SetTracerProvider(Tp)

		// Finally, set the tracer that can be used for this package.
		Tracer = Tp.Tracer("SSO")
	}
}

func newTraceProvider(exp sdkTrace.SpanExporter) *sdkTrace.TracerProvider {
	// Ensure default SDK resources and the required service name are set.
	r, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			"https://opentelemetry.io/schemas/1.26.0",
			semconv.ServiceName("SSO"),
			semconv.ServiceVersionKey.String("v0.1.0"),
			attribute.String("environment", "demo"),
		),
	)

	if err != nil {
		panic(err)
	}

	return sdkTrace.NewTracerProvider(
		sdkTrace.WithBatcher(exp),
		sdkTrace.WithResource(r),
		sdkTrace.WithSampler(sdkTrace.ParentBased(sdkTrace.AlwaysSample())),
	)
}

func newExporter(ctx context.Context) (sdkTrace.SpanExporter, error) {
	endpoint := os.Getenv("JAEGER_ENDPOINT")
	if endpoint == "" {
		endpoint = "localhost:4318"
	}

	client := otlptracehttp.NewClient(
		otlptracehttp.WithEndpoint(endpoint), //Replace Endpoint with the endpoint obtained in the Prerequisites section.
		otlptracehttp.WithInsecure())
	otlptracehttp.WithCompression(1)

	return otlptrace.New(ctx, client)
}

func StartNewTrace(ctx context.Context, name string) (context.Context, trace.Span) {
	if Tracer == nil {
		return ctx, nil
	}

	return Tracer.Start(ctx, name)
}

func EndSpan(span trace.Span) {
	if span != nil {
		span.End()
	}
}

func AddEvent(span trace.Span, name string, attributes ...trace.EventOption) {
	if span != nil {
		span.AddEvent(name, attributes...)
	}
}
