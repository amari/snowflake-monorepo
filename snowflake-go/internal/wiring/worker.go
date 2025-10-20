package wiring

import (
	"github.com/amari/snowflake-monorepo/snowflake-go/internal/config"
	"github.com/amari/snowflake-monorepo/snowflake-go/internal/temporal"
	snowflakev1 "github.com/amari/snowflake-monorepo/snowflake-go/pkg/proto/snowflake/v1"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"go.uber.org/fx"
)

func WorkerOption(cfg *config.WorkerConfig) fx.Option {
	return fx.Options(
		fx.Decorate(func(temporalClient client.Client, snowflakeClient snowflakev1.SnowflakeServiceClient, lc fx.Lifecycle) (client.Client, error) {
			w := worker.New(temporalClient, cfg.Temporal.TaskQueue, worker.Options{})

			lc.Append(fx.StartStopHook(w.Start, w.Stop))

			ao := &temporal.GRPCActivityObject{
				SnowflakeClient: snowflakeClient,
			}
			w.RegisterActivity(ao)

			return temporalClient, nil
		}),
	)
}
