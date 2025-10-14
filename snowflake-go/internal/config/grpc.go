package config

type GRPCClientConfig struct {
	Target string `koanf:"target"`

	TLS *TLSClientConfig `koanf:"tls"`
}

func DefaultGRPCClientConfig() GRPCClientConfig {
	return GRPCClientConfig{
		Target: "localhost:50051",
		TLS:    nil,
	}
}

func (c *GRPCClientConfig) Validate() error {
	return nil
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
