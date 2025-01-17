package definitions

import (
	"github.com/0x16F/cloud-users/internal/usecase/config"
	"github.com/sarulabs/di"
)

const (
	ConfigDef = "config"
)

func getConfigDef() di.Def {
	return di.Def{
		Name:  ConfigDef,
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return config.New()
		},
	}
}
