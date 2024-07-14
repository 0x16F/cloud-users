package config

import (
	"github.com/0x16F/cloud/users/internal/infrastructure/repo"
	"github.com/0x16F/cloud/users/pkg/logger"
	"github.com/ilyakaznacheev/cleanenv"
)

type WebServer struct {
	Port uint16 `env:"APP_PORT" env-default:"8080"`
}

type App struct {
	ErrorsPath     string             `env:"ERRORS_PATH"`
	MigrationsPath string             `env:"MIGRATIONS_PATH"`
	Level          logger.LoggerLevel `env:"LOGGER_LEVEL" env-default:"info"`
}

type Config struct {
	WebServer WebServer
	Database  repo.Config
	App       App
}

func New() (*Config, error) {
	cfg := Config{}

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
