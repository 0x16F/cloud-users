package definitions

import (
	"github.com/0x16F/cloud/users/internal/controller/httpsrv/handlers/users"
	"github.com/0x16F/cloud/users/internal/usecase/errors"
	usersService "github.com/0x16F/cloud/users/internal/usecase/users"
	"github.com/0x16F/cloud/users/pkg/logger"
	"github.com/sarulabs/di"
)

const (
	UsersHandlerDef = "users_handler"
)

func getUsersHandlerDef() di.Def {
	return di.Def{
		Name:  UsersHandlerDef,
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			log, _ := ctn.Get(LoggerDef).(logger.Logger)
			usersService, _ := ctn.Get(UsersServiceDef).(*usersService.Service)
			errorsService, _ := ctn.Get(ErrorsServiceDef).(errors.Errors)

			return users.NewHandler(log, usersService, errorsService), nil
		},
	}
}
