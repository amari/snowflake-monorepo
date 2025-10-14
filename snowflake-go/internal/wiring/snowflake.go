package wiring

import (
	"github.com/amari/snowflake-monorepo/snowflake-go/internal/config"
	"github.com/amari/snowflake-monorepo/snowflake-go/internal/snowflake"
	"go.uber.org/fx"
)

func SnowflakeOption(cfg *config.SnowflakeConfig) fx.Option {
	return fx.Options(
		fx.Provide(func() (*snowflake.SnowflakeService, error) {
			return snowflake.NewSnowflakeService(snowflake.SnowflakeServiceOptions{
				MachineID: uint16(cfg.MachineID),
			})
		}),
	)
}
