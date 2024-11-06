package env

import (
	"chat-server/internal/config"
	"errors"
	"net"
	"os"
)

var _ config.GRPCConfig = (*grpcConfig)(nil)

const (
	grpcHostEnvName = "GRPC_HOST"
	grpcPortEnvName = "GRPC_PORT"
)

type grpcConfig struct {
	host string
	port string
}

func NewGRPCConfig() (*grpcConfig, error) {
	host := os.Getenv(grpcHostEnvName)
	if len(host) == 0 {
		return nil, errors.New("Incorrect host config")
	}
	port := os.Getenv(grpcPortEnvName)
	if len(port) == 0 {
		return nil, errors.New("Incorrect port config")
	}
	return &grpcConfig{
		port: port,
		host: host,
	}, nil
}

func (g *grpcConfig) Address() string {
	return net.JoinHostPort(g.host, g.port)
}
