package config

import (
	"github.com/0x16F/cloud-common/pkg/logger"
	"github.com/0x16F/cloud-users/internal/infrastructure/repo"
	"github.com/ilyakaznacheev/cleanenv"
)

type App struct {
	Port           uint16             `env:"WEB_PORT" env-default:"8080"`
	Name           string             `env:"APP_NAME" env-default:"cloud-users"`
	ErrorsPath     string             `env:"ERRORS_PATH"`
	MigrationsPath string             `env:"MIGRATIONS_PATH"`
	Level          logger.LoggerLevel `env:"LOGGER_LEVEL" env-default:"info"`
	ProxyEndpoint  string             `env:"FFLAGS_ENDPOINT" env-default:"http://localhost:1031"`
}

type Config struct {
	Database repo.Config
	App      App
}

func New() (*Config, error) {
	cfg := Config{}

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
