package env

import (
	"errors"
	"net"
	"os"

	"github.com/kirillmc/trainings-server/internal/config"
)

const (
	accessHostEnvName = "ACCESS_HOST"
	accessPortEnvName = "ACCESS_PORT"
)

type accessConfig struct {
	host string
	port string
}

func NewAccessConfig() (config.AccessConfig, error) {
	host := os.Getenv(accessHostEnvName)
	if len(host) == 0 {
		return nil, errors.New("access host not found")
	}

	port := os.Getenv(accessPortEnvName)
	if len(port) == 0 {
		return nil, errors.New("access  port not found")
	}

	return &accessConfig{
		host: host,
		port: port,
	}, nil
}

func (cfg *accessConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}
