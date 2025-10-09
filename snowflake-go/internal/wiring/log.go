package wiring

import (
	"io"
	"log/slog"
	"os"
	"strings"

	fxeventzerolog "github.com/amari/fxevent-zerolog"
	"github.com/rs/zerolog"
	slogzerolog "github.com/samber/slog-zerolog/v2"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"

	"github.com/amari/snowflake-monorepo/snowflake-go/internal/config"
)

func LogOption(cfg *config.LogConfig) fx.Option {
	var writer io.Writer = os.Stderr
	if cfg.Format == "console" {
		writer = zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: "2006-01-02 15:04:05"}
	}

	lvl, err := zerolog.ParseLevel(strings.ToLower(cfg.Level))
	if err != nil {
		lvl = zerolog.InfoLevel
	}

	zerolog.SetGlobalLevel(lvl)

	// Create a new logger instance
	log := zerolog.New(writer).Level(lvl).With().Timestamp().Logger()

	stdLog := slog.New(slogzerolog.Option{
		Logger: &log,
	}.NewZerologHandler())

	slog.SetDefault(stdLog)

	return fx.Options(
		fx.WithLogger(func(logger *zerolog.Logger) fxevent.Logger {
			log := logger.With().Str("library.name", "fx").Logger()
			return fxeventzerolog.New(
				&log,
			)
		}),
		fx.Supply(&log, &stdLog),
		fx.Provide(
			func(l *zerolog.Logger) *slog.Logger {
				return slog.New(slogzerolog.Option{
					Logger: l,
				}.NewZerologHandler())
			},
		),
	)
}
