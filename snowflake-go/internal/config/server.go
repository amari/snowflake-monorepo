package config

type ServerConfig struct {
	BaseConfig `koanf:",squash"`

	GRPC      GRPCServerConfig `koanf:"grpc"`
	Snowflake SnowflakeConfig  `koanf:"snowflake"`
}

func DefaultServerConfig() ServerConfig {
	return ServerConfig{
		BaseConfig: DefaultBaseConfig(),
		GRPC:       DefaultGRPCServerConfig(),
	}
}

func (sc *ServerConfig) Validate() error {
	if err := sc.BaseConfig.Validate(); err != nil {
		return err
	}

	if err := sc.GRPC.Validate(); err != nil {
		return err
	}

	if err := sc.Snowflake.Validate(); err != nil {
		return err
	}

	return nil
}
