package wiring

import (
	"context"
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
		fx.Decorate(func(s *grpc.Server, snowflakeService *snowflake.SnowflakeService, metrics *api.SnowflakeServiceServerMetrics) *grpc.Server {
			snowflakev1.RegisterSnowflakeServiceServer(s, api.NewSnowflakeServiceServer(snowflakeService, metrics))

			return s
		}),
		fx.Provide(func(mp metric.MeterProvider) (*api.SnowflakeServiceServerMetrics, error) {
			meter := mp.Meter("snowflake-go")

			snowflakeCounter, err := meter.Int64Counter("snowflakes",
				metric.WithDescription("Total number of Snowflake IDs successfully generated"),
			)
			if err != nil {
				return nil, err
			}

			clockBackwardsErrorsCounter, err := meter.Int64Counter("snowflake_clock_backwards_errors",
				metric.WithDescription("Number of times the Snowflake service detected the system clock moved backwards"),
			)
			if err != nil {
				return nil, err
			}

			sequenceOverflowErrorsCounter, err := meter.Int64Counter("snowflake_sequence_overflow_errors",
				metric.WithDescription("Number of times the Snowflake sequence overflowed within the same millisecond"),
			)
			if err != nil {
				return nil, err
			}

			return &api.SnowflakeServiceServerMetrics{
				SnowflakeCounter:              snowflakeCounter,
				ClockBackwardsErrorsCounter:   clockBackwardsErrorsCounter,
				SequenceOverflowErrorsCounter: sequenceOverflowErrorsCounter,
			}, nil
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
				grpc.ChainUnaryInterceptor(
					grpcZerologUnaryServerInterceptor(&grpcLog),
				),
				grpc.ChainStreamInterceptor(
					grpcZerologStreamServerInterceptor(&grpcLog),
				),
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

func APIClientOption(cfg *config.GRPCClientConfig) fx.Option {
	return fx.Options(
		fx.Provide(func(logger *zerolog.Logger, lc fx.Lifecycle, sd fx.Shutdowner, meterProvider metric.MeterProvider, tracerProvider trace.TracerProvider, propogrator propagation.TextMapPropagator) (snowflakev1.SnowflakeServiceClient, error) {
			grpcLog := logger.With().
				Str(string(semconv.RPCSystemKey), semconv.RPCSystemGRPC.Value.AsString()).
				Logger()

			opts := []grpc.DialOption{
				grpc.WithStatsHandler(otelgrpc.NewClientHandler(
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
				opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)))
			} else {
				opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
			}

			conn, err := grpc.NewClient(cfg.Target, opts...)
			if err != nil {
				return nil, err
			}

			lc.Append(fx.StopHook(conn.Close))

			cli := snowflakev1.NewSnowflakeServiceClient(conn)

			grpcLog.Info().Str("target", cfg.Target).Msg("connected to snowflake service")

			return cli, nil
		}),
	)
}

func grpcZerologUnaryServerInterceptor(log *zerolog.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		newCtx := log.WithContext(ctx)
		return handler(newCtx, req)
	}
}

func grpcZerologStreamServerInterceptor(log *zerolog.Logger) grpc.StreamServerInterceptor {
	return func(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		newCtx := log.WithContext(ss.Context())
		wrapped := &grpcWrappedServerStream{ServerStream: ss, ctx: newCtx}
		return handler(srv, wrapped)
	}
}

type grpcWrappedServerStream struct {
	grpc.ServerStream
	ctx context.Context
}

func (w *grpcWrappedServerStream) Context() context.Context {
	return w.ctx
}
