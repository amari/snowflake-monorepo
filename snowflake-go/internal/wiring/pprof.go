package wiring

import (
	"context"
	"net"
	"net/http"
	"net/http/pprof"
	"strconv"
	"time"

	"github.com/amari/snowflake-monorepo/snowflake-go/internal/config"
	"github.com/rs/zerolog"
	semconv "go.opentelemetry.io/otel/semconv/v1.37.0"
	"go.uber.org/fx"
)

func PprofOption(cfg *config.PProfConfig) fx.Option {
	if !cfg.Enabled {
		return fx.Options()
	}

	return fx.Options(
		fx.Invoke(func(lc fx.Lifecycle, sd fx.Shutdowner, logger *zerolog.Logger) error {
			log := logger.With().Str("component", "pprof").Logger()

			// Set up pprof handlers
			mux := http.NewServeMux()
			mux.HandleFunc("/debug/pprof/", pprof.Index)
			mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
			mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
			mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
			mux.HandleFunc("/debug/pprof/trace", pprof.Trace)

			// Create the HTTP server
			s := &http.Server{
				Handler: mux,
				// Reasonable production defaults
				ReadTimeout:       15 * time.Second,
				ReadHeaderTimeout: 5 * time.Second,
				WriteTimeout:      30 * time.Second,
				IdleTimeout:       120 * time.Second,
			}

			lc.Append(fx.StartStopHook(func(ctx context.Context) error {
				// Listen on the configured address and port
				hostPort := net.JoinHostPort(cfg.Address, strconv.FormatUint(uint64(cfg.Port), 10))

				lis, err := net.Listen("tcp", hostPort)
				if err != nil {
					return err
				}

				log.Info().
					Str(string(semconv.ServerAddressKey), lis.Addr().String()).
					Msg("pprof server listening")

				// Start pprof server
				go func() {
					if err := s.Serve(lis); err != nil && err != http.ErrServerClosed {
						log.Error().Err(err).Msg("pprof server stopped unexpectedly")

						_ = sd.Shutdown()
					}
				}()

				return nil
			}, s.Shutdown))

			return nil
		}),
	)
}
