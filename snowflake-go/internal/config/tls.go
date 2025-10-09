package config

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"os"
)

type TLSClientConfig struct {
	// CAFile is the path to a PEM-encoded file containing one or more trusted
	// root or intermediate certificates. These are used to verify the server’s
	// certificate chain when CertPoolName is not set.
	//
	// This field configures *trust*, not identity.
	CAFile string `koanf:"caFile"`

	// CertFile is the path to the client’s own certificate (PEM-encoded). If set,
	// it will be presented to the server during the TLS handshake for client
	// authentication (mTLS).
	//
	// Optional — omit if client authentication is not required.
	CertFile string `koanf:"certFile"`

	// KeyFile is the path to the private key (PEM-encoded) corresponding to CertFile.
	// Required if CertFile is set.
	//
	// This key should be protected and never shared between components.
	KeyFile string `koanf:"keyFile"`
}

func (cc *TLSClientConfig) Validate() error {
	errs := []error{}

	if !(cc.CertFile != "" && cc.KeyFile != "") {
		errs = append(errs, fmt.Errorf("tls: both certFile and keyFile must be set"))
	}

	if errs != nil {
		return errors.Join(errs...)
	}

	return nil
}

func (cfg *TLSClientConfig) ToTLSConfig() (*tls.Config, error) {
	// Loads a client identity certificate (cert+key pair) if both fields are set.
	loadClientIdentity := func(certFile, keyFile string) ([]tls.Certificate, error) {
		if certFile == "" && keyFile == "" {
			return nil, nil
		}
		if certFile == "" || keyFile == "" {
			return nil, fmt.Errorf("tls both certFile and keyFile must be set")
		}
		cert, err := tls.LoadX509KeyPair(certFile, keyFile)
		if err != nil {
			return nil, fmt.Errorf("tls failed to load cert/key: %w", err)
		}
		return []tls.Certificate{cert}, nil
	}

	var pool *x509.CertPool
	var err error

	switch {
	case cfg.CAFile != "":
		pool = x509.NewCertPool()
		pem, err := os.ReadFile(cfg.CAFile)
		if err != nil {
			return nil, fmt.Errorf("tls failed to read CA file %q: %w", cfg.CAFile, err)
		}
		if !pool.AppendCertsFromPEM(pem) {
			return nil, fmt.Errorf("tls invalid PEM in CA file %q", cfg.CAFile)
		}
	default:
		pool, err = x509.SystemCertPool()
		if err != nil {
			return nil, fmt.Errorf("tls failed to load system cert pool: %w", err)
		}
	}

	tlsCfg := &tls.Config{
		RootCAs: pool,
	}

	certs, err := loadClientIdentity(cfg.CertFile, cfg.KeyFile)
	if err != nil {
		return nil, err
	}
	tlsCfg.Certificates = certs

	return tlsCfg, nil
}

type TLSServerConfig struct {
	// CAFile is the path to a PEM-encoded file containing trusted root/intermediate
	// certificates used to verify client certificates. Used only if CertPoolName is not set.
	//
	// This does *not* affect the server’s own identity — it controls whether and how
	// client certificates are verified.
	CAFile string `koanf:"caFile"`

	// CertFile is the path to the server’s certificate (PEM-encoded). This is the
	// public certificate presented to clients during the TLS handshake.
	//
	// Required.
	CertFile string `koanf:"certFile"`

	// KeyFile is the path to the server’s private key (PEM-encoded) corresponding
	// to CertFile.
	//
	// Required. Should be protected and not shared.
	KeyFile string `koanf:"keyFile"`
}

func (cc *TLSServerConfig) Validate() error {
	errs := []error{}

	if !(cc.CertFile != "" && cc.KeyFile != "") {
		errs = append(errs, fmt.Errorf("tls: both certFile and keyFile must be set"))
	}

	if errs != nil {
		return errors.Join(errs...)
	}

	return nil
}

func (cfg *TLSServerConfig) ToTLSConfig() (*tls.Config, error) {
	// Loads a server identity certificate (cert+key pair) if both fields are set.
	loadServerIdentity := func(certFile, keyFile string) ([]tls.Certificate, error) {
		if certFile == "" && keyFile == "" {
			return nil, nil
		}
		if certFile == "" || keyFile == "" {
			return nil, fmt.Errorf("tls both certFile and keyFile must be set")
		}
		cert, err := tls.LoadX509KeyPair(certFile, keyFile)
		if err != nil {
			return nil, fmt.Errorf("tls failed to load cert/key: %w", err)
		}
		return []tls.Certificate{cert}, nil
	}

	certs, err := loadServerIdentity(cfg.CertFile, cfg.KeyFile)
	if err != nil {
		return nil, err
	}

	tlsCfg := &tls.Config{
		Certificates: certs,
	}

	var pool *x509.CertPool

	switch {
	case cfg.CAFile != "":
		pool = x509.NewCertPool()
		pem, err := os.ReadFile(cfg.CAFile)
		if err != nil {
			return nil, fmt.Errorf("tls failed to read CA file %q: %w", cfg.CAFile, err)
		}
		if !pool.AppendCertsFromPEM(pem) {
			return nil, fmt.Errorf("tls invalid PEM in CA file %q", cfg.CAFile)
		}
		tlsCfg.ClientCAs = pool
		tlsCfg.ClientAuth = tls.RequireAndVerifyClientCert
	default:
		tlsCfg.ClientAuth = tls.NoClientCert
	}

	return tlsCfg, nil
}
