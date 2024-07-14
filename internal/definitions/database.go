package definitions

import (
	"context"

	"github.com/0x16F/cloud/users/internal/infrastructure/repo"
	"github.com/0x16F/cloud/users/internal/usecase/config"
	"github.com/0x16F/cloud/users/internal/usecase/migrations"
	"github.com/sarulabs/di"
)

const (
	DatabaseDef = "database"
)

func getDatabaseDef() di.Def {
	return di.Def{
		Name:  DatabaseDef,
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			ctx, _ := ctn.Get(ContextDef).(context.Context)
			cfg, _ := ctn.Get(ConfigDef).(*config.Config)

			pool, err := repo.NewConnection(ctx, cfg.Database)
			if err != nil {
				return nil, err
			}

			if err := migrations.Up(pool, cfg.App.MigrationsPath); err != nil {
				return nil, err
			}

			return pool, nil
		},
	}
}
