package tracing

import (
	"net/http"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/zipkin"
)

// Inject injects the outbound HTTP request with the given span's context to ensure
// correct propagation of span context throughout the trace.
func Inject(span opentracing.Span, request *http.Request) error {
	return span.Tracer().Inject(
		span.Context(),
		opentracing.HTTPHeaders,
		zipkin.NewZipkinB3HTTPHeaderPropagator())
}

// Extract extracts the inbound HTTP request to obtain the parent span's context to ensure
// correct propagation of span context throughout the trace.
func Extract(tracer opentracing.Tracer, r *http.Request) (opentracing.SpanContext, error) {
	return tracer.Extract(
		opentracing.HTTPHeaders,
		zipkin.NewZipkinB3HTTPHeaderPropagator())
}
