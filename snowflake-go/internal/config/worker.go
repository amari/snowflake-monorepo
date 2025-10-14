package config

type WorkerConfig struct {
	BaseConfig `koanf:",squash"`

	GRPC     GRPCClientConfig     `koanf:"grpc"`
	Temporal TemporalClientConfig `koanf:"temporal"`
}

func DefaultWorkerConfig() WorkerConfig {
	return WorkerConfig{
		BaseConfig: DefaultBaseConfig(),
		GRPC:       DefaultGRPCClientConfig(),
		Temporal:   DefaultTemporalClientConfig(),
	}
}

func (sc *WorkerConfig) Validate() error {
	if err := sc.BaseConfig.Validate(); err != nil {
		return err
	}

	if err := sc.GRPC.Validate(); err != nil {
		return err
	}

	return nil
}
