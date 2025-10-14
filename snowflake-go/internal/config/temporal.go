package config

type TemporalClientConfig struct {
	Target    string `koanf:"target"`
	Namespace string `koanf:"namespace"`
	TaskQueue string `koanf:"taskQueue"`
}

func DefaultTemporalClientConfig() TemporalClientConfig {
	return TemporalClientConfig{
		Target:    "localhost:7233",
		Namespace: "default",
		TaskQueue: "default",
	}
}
