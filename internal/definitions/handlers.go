package definitions

import (
	"github.com/0x16F/cloud-common/pkg/logger"
	"github.com/0x16F/cloud-users/internal/controller/httpsrv/features"
	"github.com/0x16F/cloud-users/internal/controller/httpsrv/handlers/users"
	"github.com/0x16F/cloud-users/internal/usecase/errors"
	"github.com/0x16F/cloud-users/internal/usecase/fflags"
	usersService "github.com/0x16F/cloud-users/internal/usecase/users"
	"github.com/sarulabs/di"
)

const (
	UsersHandlerDef    = "users_handler"
	FeaturesServiceDef = "features_service"
)

func getUsersHandlerDef() di.Def {
	return di.Def{
		Name:  UsersHandlerDef,
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			log, _ := ctn.Get(LoggerDef).(logger.Logger)
			usersService, _ := ctn.Get(UsersServiceDef).(*usersService.Service)
			errorsService, _ := ctn.Get(ErrorsServiceDef).(errors.Errors)
			featuresService, _ := ctn.Get(FeaturesServiceDef).(*features.Service)

			return users.NewHandler(log, usersService, errorsService, featuresService), nil
		},
	}
}

func getFeaturesServiceDef() di.Def {
	return di.Def{
		Name:  FeaturesServiceDef,
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			fflagsService, _ := ctn.Get(FFlagsServiceDef).(*fflags.Service)
			errorsService, _ := ctn.Get(ErrorsServiceDef).(errors.Errors)

			return features.New(fflagsService, errorsService), nil
		},
	}
}
