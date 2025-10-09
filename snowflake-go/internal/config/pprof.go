package config

type PProfConfig struct {
	Enabled bool   `koanf:"enabled"`
	Address string `koanf:"address"`
	Port    int    `koanf:"port"`
}

func DefaultPProfConfig() PProfConfig {
	return PProfConfig{
		Enabled: false,
		Address: "localhost",
		Port:    6060,
	}
}

func (pc *PProfConfig) Validate() error {
	return nil
}
