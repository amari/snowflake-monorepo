package wiring

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	exportersprometheus "go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/propagation"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.37.0"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/fx"
)

func OtelOption() fx.Option {
	return fx.Options(
		fx.Invoke(func(_ metric.MeterProvider, _ trace.TracerProvider) {}),
		OtelResourceOption(),
		fx.Provide(
			func(r *resource.Resource, lc fx.Lifecycle, sd fx.Shutdowner, logger *zerolog.Logger) (metric.MeterProvider, error) {
				log := logger.With().Str("component", "metrics").Logger()

				exporterType := strings.ToLower(os.Getenv("OTEL_METRICS_EXPORTER"))
				if exporterType == "" {
					exporterType = "none"
				}

				var reader sdkmetric.Reader

				switch exporterType {
				case "none":
					// No exporter
				case "prometheus":
					exporterOpts := []exportersprometheus.Option{
						// exportersprometheus.WithoutUnits(),
						// exportersprometheus.WithoutScopeInfo(),
						// exportersprometheus.WithAggregationSelector(defaultAggregationSelector),
						// exportersprometheus.WithRegisterer(prometheus.DefaultRegisterer),
					}

					exporter, err := exportersprometheus.New(exporterOpts...)
					if err != nil {
						return nil, fmt.Errorf("failed to create Prometheus exporter: %w", err)
					}

					reader = exporter

				default:
					return nil, fmt.Errorf("unsupported OTEL_METRICS_EXPORTER: %q", exporterType)
				}

				metricProvider := sdkmetric.NewMeterProvider(
					sdkmetric.WithResource(r),
					sdkmetric.WithReader(reader),
				)

				lc.Append(fx.Hook{
					OnStop: func(ctx context.Context) error {
						return metricProvider.Shutdown(ctx)
					},
				})

				if exporterType == "prometheus" {
					host := os.Getenv("OTEL_EXPORTER_PROMETHEUS_HOST")

					port := os.Getenv("OTEL_EXPORTER_PROMETHEUS_PORT")
					if port == "" {
						port = "9464"
					} else if _, err := strconv.ParseUint(port, 10, 16); err != nil {
						return nil, fmt.Errorf("unsupported OTEL_EXPORTER_PROMETHEUS_PORT: %q", port)
					}

					addr := net.JoinHostPort(host, port)

					mux := http.NewServeMux()

					mux.Handle("/metrics", promhttp.Handler())

					srv := &http.Server{
						Handler:      mux,
						ReadTimeout:  5 * time.Second,
						WriteTimeout: 10 * time.Second,
						IdleTimeout:  120 * time.Second,
					}

					lc.Append(fx.StartStopHook(func(ctx context.Context) error {
						lis, err := net.Listen("tcp", addr)
						if err != nil {
							return fmt.Errorf("failed to create Prometheus /metrics endpoint: %w", err)
						}

						log.Info().
							Str(string(semconv.ServerAddressKey), lis.Addr().String()).
							Msg("prometheus /metrics endpoint listening")

						go func() {
							defer lis.Close()

							if err := srv.Serve(lis); err != http.ErrServerClosed {
								log.Error().Err(err).Msg("prometheus /metrics endpoint stopped unexpectedly")

								_ = sd.Shutdown()
							}
						}()

						return nil
					}, srv.Shutdown))
				}

				// Set global providers
				otel.SetMeterProvider(metricProvider)

				return metricProvider, nil
			},
		),
		fx.Provide(
			func(tp *sdktrace.TracerProvider) trace.TracerProvider { return tp },
			func(r *resource.Resource, lc fx.Lifecycle, logger *zerolog.Logger) (*sdktrace.TracerProvider, error) {
				traceExporterCtx := context.Background()

				traceExporter, err := otelTraceExporterFromEnv(traceExporterCtx)
				if err != nil {
					return nil, err
				}

				traceSampler := otelTraceSamplerFromEnv(logger)

				traceProvider := sdktrace.NewTracerProvider(
					sdktrace.WithBatcher(traceExporter),
					sdktrace.WithResource(r),
					sdktrace.WithSampler(traceSampler),
				)

				lc.Append(fx.Hook{
					OnStop: func(ctx context.Context) error {
						return traceProvider.Shutdown(ctx)
					},
				})

				// Set global providers
				otel.SetTracerProvider(traceProvider)

				return traceProvider, nil
			},
		),
		fx.Provide(
			func() propagation.TextMapPropagator {
				p := otelTracePropagatorFromEnv()

				otel.SetTextMapPropagator(p)

				return p
			},
		),
	)
}

func OtelResourceOption() fx.Option {
	return fx.Options(
		fx.Provide(func() (*resource.Resource, error) {
			return otelResourceFromEnv(context.Background())
		}),
	)
}

func otelResourceFromEnv(ctx context.Context) (*resource.Resource, error) {
	var attrs []attribute.KeyValue

	// Standard service fields
	if serviceName := os.Getenv("OTEL_SERVICE_NAME"); serviceName != "" {
		attrs = append(attrs, semconv.ServiceName(serviceName))
	}
	if serviceVersion := os.Getenv("OTEL_SERVICE_VERSION"); serviceVersion != "" {
		attrs = append(attrs, semconv.ServiceVersion(serviceVersion))
	}
	if serviceNamespace := os.Getenv("OTEL_SERVICE_NAMESPACE"); serviceNamespace != "" {
		attrs = append(attrs, semconv.ServiceNamespace(serviceNamespace))
	}

	// OTEL_RESOURCE_ATTRIBUTES parsing
	if rawAttrs := os.Getenv("OTEL_RESOURCE_ATTRIBUTES"); rawAttrs != "" {
		pairs := strings.Split(rawAttrs, ",")
		for _, pair := range pairs {
			kv := strings.SplitN(pair, "=", 2)
			if len(kv) == 2 {
				key := strings.TrimSpace(kv[0])
				value := strings.TrimSpace(kv[1])
				attrs = append(attrs, attribute.String(key, value))
			}
		}
	}

	return resource.New(ctx,
		resource.WithAttributes(attrs...),
		resource.WithFromEnv(),      // Also allow OTEL_RESOURCE_ATTRIBUTES
		resource.WithTelemetrySDK(), // Add SDK information
		resource.WithHost(),         // Add host information
	)
}

func otelTraceExporterFromEnv(ctx context.Context) (sdktrace.SpanExporter, error) {
	protocol := os.Getenv("OTEL_EXPORTER_OTLP_PROTOCOL")
	endpoint := os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT")
	if endpoint == "" {
		endpoint = "localhost:4317" // default according to OpenTelemetry spec
	}

	// Build exporter based on protocol

	switch protocol {
	case "grpc", "":
		// Default is gRPC
		return otlptracegrpc.New(ctx,
			otlptracegrpc.WithEndpoint(endpoint),
			otlptracegrpc.WithInsecure(),
		)
	case "http/protobuf":
		return otlptracehttp.New(ctx,
			otlptracehttp.WithEndpoint(endpoint),
			otlptracehttp.WithInsecure(),
		)
	default:
		return nil, fmt.Errorf("unsupported OTEL_EXPORTER_OTLP_PROTOCOL: %s", protocol)
	}
}

// otelTraceSamplerFromEnv builds a trace.Sampler based on OTEL_TRACES_SAMPLER and OTEL_TRACES_SAMPLER_ARG.
func otelTraceSamplerFromEnv(log *zerolog.Logger) sdktrace.Sampler {
	name := os.Getenv("OTEL_TRACES_SAMPLER")
	arg := os.Getenv("OTEL_TRACES_SAMPLER_ARG")

	switch strings.ToLower(name) {
	case "always_on":
		return sdktrace.AlwaysSample()
	case "always_off":
		return sdktrace.NeverSample()
	case "traceidratio":
		ratio := otelTraceParseRatio(log, arg)
		return sdktrace.TraceIDRatioBased(ratio)
	case "parentbased_always_on":
		return sdktrace.ParentBased(sdktrace.AlwaysSample())
	case "parentbased_always_off":
		return sdktrace.ParentBased(sdktrace.NeverSample())
	case "parentbased_traceidratio":
		ratio := otelTraceParseRatio(log, arg)
		return sdktrace.ParentBased(sdktrace.TraceIDRatioBased(ratio))
	case "", "default":
		return sdktrace.ParentBased(sdktrace.TraceIDRatioBased(1.0))
	default:
		// fallback to always_on with warning
		log.Warn().Str("system", "otel").Msgf("unknown OTEL_TRACES_SAMPLER %q, falling back to AlwaysSample", name)

		return sdktrace.AlwaysSample()
	}
}

func otelTraceParseRatio(log *zerolog.Logger, s string) float64 {
	if s == "" {
		return 1.0
	}
	ratio, err := strconv.ParseFloat(s, 64)
	if err != nil || ratio < 0 || ratio > 1 {
		log.Warn().Str("system", "otel").Msgf("invalid OTEL_TRACES_SAMPLER_ARG %q, using default ratio 1.0", s)

		return 1.0
	}
	return ratio
}

func otelTracePropagatorFromEnv() propagation.TextMapPropagator {
	env := os.Getenv("OTEL_PROPAGATORS")
	if env == "" {
		env = "tracecontext,baggage" // default according to OpenTelemetry spec
	}

	propagators := []propagation.TextMapPropagator{}

	// TODO: add support for b3, b3multi, xray, etc.
	for _, p := range strings.Split(env, ",") {
		switch strings.TrimSpace(strings.ToLower(p)) {
		case "tracecontext":
			propagators = append(propagators, propagation.TraceContext{})
		case "baggage":
			propagators = append(propagators, propagation.Baggage{})
		}
	}

	if len(propagators) == 0 {
		return propagation.NewCompositeTextMapPropagator()
	}

	return propagation.NewCompositeTextMapPropagator(propagators...)
}
