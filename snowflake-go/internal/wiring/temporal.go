package wiring

import (
	"github.com/amari/snowflake-monorepo/snowflake-go/internal/config"
	"github.com/rs/zerolog"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/contrib/opentelemetry"
	"go.temporal.io/sdk/interceptor"
	"go.temporal.io/sdk/log"
	"go.uber.org/fx"
)

func TemporalClientOption(cfg *config.TemporalClientConfig) fx.Option {
	return fx.Options(
		fx.Invoke(func(temporalClient client.Client, lc fx.Lifecycle, sd fx.Shutdowner) {
			lc.Append(fx.StopHook(temporalClient.Close))
		}),
		fx.Provide(func(meterProvider metric.MeterProvider, tracerProvider trace.TracerProvider, propogrator propagation.TextMapPropagator, logger *zerolog.Logger) (client.Client, error) {
			tracingInterceptor, err := opentelemetry.NewTracingInterceptor(opentelemetry.TracerOptions{})
			if err != nil {
				return nil, err
			}

			metricsHandler := opentelemetry.NewMetricsHandler(opentelemetry.MetricsHandlerOptions{})

			log := logger.With().Str("component", "temporal-client").Logger()

			cli, err := client.NewLazyClient(client.Options{
				HostPort:       cfg.Target,
				Namespace:      cfg.Namespace,
				Logger:         &temporalLoggerAdapter{logger: &log},
				MetricsHandler: metricsHandler,
				Interceptors: []interceptor.ClientInterceptor{
					tracingInterceptor,
				},
			})
			if err != nil {
				return nil, err
			}

			return cli, nil
		}),
	)
}

type temporalLoggerAdapter struct {
	logger *zerolog.Logger
}

func (t *temporalLoggerAdapter) Debug(msg string, keyvals ...any) {
	t.logger.Debug().Fields(t.toFields(keyvals...)).Msg(msg)
}

func (t *temporalLoggerAdapter) Info(msg string, keyvals ...any) {
	t.logger.Info().Fields(t.toFields(keyvals...)).Msg(msg)
}

func (t *temporalLoggerAdapter) Warn(msg string, keyvals ...any) {
	t.logger.Warn().Fields(t.toFields(keyvals...)).Msg(msg)
}

func (t *temporalLoggerAdapter) Error(msg string, keyvals ...any) {
	t.logger.Error().Fields(t.toFields(keyvals...)).Msg(msg)
}

func (t *temporalLoggerAdapter) WithCallerSkip(skip int) log.Logger {
	l := t.logger.With().CallerWithSkipFrameCount(skip).Logger()
	return &temporalLoggerAdapter{logger: &l}
}

func (t *temporalLoggerAdapter) With(keyvals ...any) log.Logger {
	l := t.logger.With().Fields(t.toFields(keyvals...)).Logger()
	return &temporalLoggerAdapter{logger: &l}
}

// toFields is now a private method on temporalLoggerAdapter
func (t *temporalLoggerAdapter) toFields(keyvals ...any) map[string]any {
	fields := make(map[string]any)
	for i := 0; i < len(keyvals)-1; i += 2 {
		key, ok := keyvals[i].(string)
		if !ok {
			continue // skip non-string keys
		}
		fields[key] = keyvals[i+1]
	}
	return fields
}
