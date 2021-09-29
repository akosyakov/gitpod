package config

import (
	"crypto/tls"
	"crypto/x509"
	"github.com/gitpod-io/gitpod/content-service/api/config"
	"golang.org/x/xerrors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"os"
)

type Service struct {
	Addr string    `json:"address"`
	TLS  TLSConfig `json:"tls"`
}

type Prometheus struct {
	Addr string `json:"address"`
}

type PProf struct {
	Addr string `json:"address"`
}

type ServiceConfig struct {
	// Daemon  daemon.Config `json:"daemon"`
	Service    Service              `json:"service"`
	Prometheus Prometheus           `json:"prometheus"`
	PProf      PProf                `json:"pprof"`
	Storage    config.StorageConfig `json:"storage"`
}

type TLSConfig struct {
	Authority   string `json:"ca"`
	Certificate string `json:"crt"`
	PrivateKey  string `json:"key"`
}

// ServerOption produces the GRPC option that configures a server to use this TLS configuration
func (c *TLSConfig) ServerOption() (grpc.ServerOption, error) {
	if c.Authority == "" || c.Certificate == "" || c.PrivateKey == "" {
		return nil, nil
	}

	// Load certs
	certificate, err := tls.LoadX509KeyPair(c.Certificate, c.PrivateKey)
	if err != nil {
		return nil, xerrors.Errorf("cannot load TLS certificate: %w", err)
	}

	// Create a certificate pool from the certificate authority
	certPool := x509.NewCertPool()
	ca, err := os.ReadFile(c.Authority)
	if err != nil {
		return nil, xerrors.Errorf("cannot not read ca certificate: %w", err)
	}
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		return nil, xerrors.Errorf("failed to append ca certs")
	}

	creds := credentials.NewTLS(&tls.Config{
		ClientAuth:   tls.RequireAndVerifyClientCert,
		Certificates: []tls.Certificate{certificate},
		ClientCAs:    certPool,
		MinVersion:   tls.VersionTLS12,
	})

	return grpc.Creds(creds), nil
}
