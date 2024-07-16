package definitions

import (
	"github.com/0x16F/cloud-common/pkg/logger"
	"github.com/0x16F/cloud-users/internal/usecase/config"
	"github.com/sarulabs/di"
)

const (
	LoggerDef = "logger"
)

func getLoggerDef() di.Def {
	return di.Def{
		Name:  LoggerDef,
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg, _ := ctn.Get(ConfigDef).(*config.Config)

			return logger.New(cfg.App.Level), nil
		},
	}
}
