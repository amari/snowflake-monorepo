package wiring

import (
	"net"
	"strconv"

	"github.com/rs/zerolog"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/propagation"
	semconv "go.opentelemetry.io/otel/semconv/v1.37.0"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"

	"github.com/amari/snowflake-monorepo/snowflake-go/internal/api"
	"github.com/amari/snowflake-monorepo/snowflake-go/internal/config"
	snowflakev1 "github.com/amari/snowflake-monorepo/snowflake-go/internal/proto/snowflake/v1"
	"github.com/amari/snowflake-monorepo/snowflake-go/internal/snowflake"
)

func APIOption(cfg *config.GRPCServerConfig) fx.Option {
	return fx.Options(
		fx.Invoke(func(*grpc.Server) {}),
		fx.Decorate(func(s *grpc.Server, snowflakeService *snowflake.SnowflakeService) *grpc.Server {
			snowflakev1.RegisterSnowflakeServiceServer(s, api.NewSnowflakeServiceServer(snowflakeService))

			return s
		}),
		fx.Provide(func(logger *zerolog.Logger, lc fx.Lifecycle, sd fx.Shutdowner, meterProvider metric.MeterProvider, tracerProvider trace.TracerProvider, propogrator propagation.TextMapPropagator) (*grpc.Server, error) {
			grpcLog := logger.With().
				Str(string(semconv.RPCSystemKey), semconv.RPCSystemGRPC.Value.AsString()).
				Logger()

			opts := []grpc.ServerOption{
				grpc.StatsHandler(otelgrpc.NewServerHandler(
					otelgrpc.WithMeterProvider(meterProvider),
					otelgrpc.WithPropagators(propogrator),
					otelgrpc.WithTracerProvider(tracerProvider),
				)),
			}

			if cfg.TLS != nil {
				tlsConfig, err := cfg.TLS.ToTLSConfig()
				if err != nil {
					return nil, err
				}
				opts = append(opts, grpc.Creds(credentials.NewTLS(tlsConfig)))
			} else {
				opts = append(opts, grpc.Creds(insecure.NewCredentials()))
			}

			s := grpc.NewServer(opts...)

			reflection.Register(s)

			lc.Append(fx.StartStopHook(func() error {
				// listen and serve
				hostPort := net.JoinHostPort(cfg.Address, strconv.FormatUint(uint64(cfg.Port), 10))

				lis, err := net.Listen("tcp", hostPort)
				if err != nil {
					return err
				}

				grpcLog.Info().
					Str(string(semconv.ServerAddressKey), lis.Addr().String()).
					Msg("gRPC server listening")

				go func() {
					if err := s.Serve(lis); err != nil {
						if err != grpc.ErrServerStopped {
							grpcLog.Error().Err(err).Msg("gRPC server stopped unexpectedly")

							sd.Shutdown()
						}
					}
				}()

				return nil
			}, s.Stop))

			return s, nil
		}),
	)
}
