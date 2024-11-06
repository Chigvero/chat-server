package env

import (
	"chat-server/internal/config"
	"errors"
	"os"
)

const (
	pgDsnEnvName = "MIGRATION_DSN_L"
)

var _ config.PGConfig = (*pgConfig)(nil)

type pgConfig struct {
	dsn string
}

func NewPgConfig() (*pgConfig, error) {
	dsn := os.Getenv(pgDsnEnvName)
	if len(dsn) == 0 {
		return nil, errors.New("Incorrect dsn string")
	}
	return &pgConfig{
		dsn: dsn,
	}, nil
}

func (p *pgConfig) DSN() string {
	return p.dsn
}
