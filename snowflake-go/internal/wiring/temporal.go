package wiring

import (
	"github.com/amari/snowflake-monorepo/snowflake-go/internal/config"
	"go.uber.org/fx"
)

func TemporalClientOption(cfg *config.TemporalClientConfig) fx.Option {
	return fx.Options()
}

func TemporalWorkerOption(cfg *config.TemporalWorkerConfig) fx.Option {
	return fx.Options(
		TemporalClientOption(&cfg.Client),
	)
}
