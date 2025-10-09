package config

type BaseConfig struct {
	Log   LogConfig   `koanf:"log"`
	PProf PProfConfig `koanf:"pprof"`
}

func DefaultBaseConfig() BaseConfig {
	return BaseConfig{
		Log: LogConfig{
			Level:  "info",
			Format: "json",
		},
		PProf: DefaultPProfConfig(),
	}
}

func (bc *BaseConfig) Validate() error {
	if err := bc.Log.Validate(); err != nil {
		return err
	}

	if err := bc.PProf.Validate(); err != nil {
		return err
	}

	return nil
}
