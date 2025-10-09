package config

type TemporalWorkerConfig struct {
	BaseConfig `koanf:",squash"`

	Client TemporalClientConfig `koanf:"client"`
}

type TemporalClientConfig struct{}
