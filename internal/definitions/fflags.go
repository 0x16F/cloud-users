package definitions

import (
	"github.com/0x16F/cloud-common/pkg/fflags"
	"github.com/0x16F/cloud-common/pkg/logger"
	"github.com/0x16F/cloud-users/internal/usecase/config"
	gofeatureflag "github.com/open-feature/go-sdk-contrib/providers/go-feature-flag/pkg"
	"github.com/sarulabs/di"
)

const (
	FFlagsClientDef   = "fflags_client"
	FFlagsProviderDef = "fflags_provider"
)

func getFFlagsProviderDef() di.Def {
	return di.Def{
		Name:  FFlagsProviderDef,
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg, _ := ctn.Get(ConfigDef).(*config.Config)
			log, _ := ctn.Get(LoggerDef).(logger.Logger)

			return fflags.NewProvider(log, cfg.App.ProxyEndpoint)
		},
		Close: func(obj interface{}) error {
			obj.(*gofeatureflag.Provider).Shutdown()
			return nil
		},
	}
}

func getFFlagsClientDef() di.Def {
	return di.Def{
		Name:  FFlagsClientDef,
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg, _ := ctn.Get(ConfigDef).(*config.Config)
			log, _ := ctn.Get(LoggerDef).(logger.Logger)
			provider, _ := ctn.Get(FFlagsProviderDef).(*gofeatureflag.Provider)

			return fflags.NewClient(log, provider, cfg.App.Name)
		},
	}
}
