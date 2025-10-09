package config

type GRPCClientConfig struct {
	TLS *TLSClientConfig `koanf:"tls"`
}

type GRPCServerConfig struct {
	Address string `koanf:"addr"`
	Port    uint16 `koanf:"port"`

	TLS *TLSServerConfig `koanf:"tls"`
}

func DefaultGRPCServerConfig() GRPCServerConfig {
	return GRPCServerConfig{
		Address: "localhost",
		Port:    50051,
		TLS:     nil,
	}
}

func (c *GRPCServerConfig) Validate() error {
	return nil
}
