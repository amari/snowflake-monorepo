package main

import (
	"context"
	"log"
	"os"

	"github.com/amari/snowflake-monorepo/snowflake-go/internal/config"
	"github.com/amari/snowflake-monorepo/snowflake-go/internal/wiring"
	"github.com/knadh/koanf/v2"
	"github.com/urfave/cli/v3"
	_ "go.uber.org/automaxprocs"
	"go.uber.org/fx"
)

func main() {
	cmd := &cli.Command{
		Name:   "temporal-worker",
		Usage:  "A temporal worker for generating snowflakes",
		Action: Action,
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

func Action(ctx context.Context, cmd *cli.Command) error {
	// Create a new Koanf instance
	k := koanf.New(".")

	// Load from environment variables with the "SNOWFLAKE_" prefix
	if err := config.LoadConfigEnvVars(k, "SNOWFLAKE_", "."); err != nil {
		return err
	}

	// Unmarshal the Config
	cfg := config.DefaultWorkerConfig()
	if err := k.Unmarshal("", &cfg); err != nil {
		return err
	}

	// Validate the Config
	if err := cfg.Validate(); err != nil {
		return err
	}

	// Start the worker
	fx.New(
		wiring.APIClientOption(&cfg.GRPC),
		wiring.LogOption(&cfg.Log),
		wiring.OtelOption(),
		wiring.PprofOption(&cfg.PProf),
		wiring.TemporalClientOption(&cfg.Temporal),
		wiring.WorkerOption(&cfg),
	).Run()

	return nil
}
