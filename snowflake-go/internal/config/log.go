package config

import (
	"errors"
	"strings"

	"github.com/rs/zerolog"
)

type LogConfig struct {
	Level string `koanf:"level"`
	// Format is either "json" or "console"
	Format string `koanf:"format"`
}

func (lc *LogConfig) Validate() error {
	if lc.Level == "" {
		lc.Level = "info"
	}

	if lc.Format == "" {
		lc.Format = "console"
	}

	if lc.Format != "json" && lc.Format != "console" {
		return errors.New("invalid log format")
	}

	_, err := zerolog.ParseLevel(strings.ToLower(lc.Level))

	return err
}
