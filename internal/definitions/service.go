package definitions

import (
	"github.com/0x16F/cloud/users/internal/infrastructure/repo/users"
	"github.com/0x16F/cloud/users/internal/usecase/config"
	"github.com/0x16F/cloud/users/internal/usecase/errors"
	usersService "github.com/0x16F/cloud/users/internal/usecase/users"
	"github.com/0x16F/cloud/users/pkg/logger"
	"github.com/sarulabs/di"
)

const (
	UsersServiceDef  = "users_service"
	ErrorsServiceDef = "errors_service"
)

func getUsersServiceDef() di.Def {
	return di.Def{
		Name:  UsersServiceDef,
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			log, _ := ctn.Get(LoggerDef).(logger.Logger)
			usersRepo, _ := ctn.Get(UsersRepoDef).(*users.Repo)
			errorsService, _ := ctn.Get(ErrorsServiceDef).(errors.Errors)

			return usersService.NewService(log, usersRepo, errorsService), nil
		},
	}
}

func getErrorsServiceDef() di.Def {
	return di.Def{
		Name:  ErrorsServiceDef,
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			log, _ := ctn.Get(LoggerDef).(logger.Logger)
			cfg, _ := ctn.Get(ConfigDef).(*config.Config)

			return errors.New(log, cfg.App.ErrorsPath), nil
		},
	}
}
