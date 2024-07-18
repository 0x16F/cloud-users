package definitions

import (
	"context"

	"github.com/0x16F/cloud-users/internal/infrastructure/repo/users"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sarulabs/di/v2"
)

const (
	UsersRepoDef = "users_repo"
)

func getUsersRepoDef() *di.Def {
	return &di.Def{
		Name:  UsersRepoDef,
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			pool, _ := ctn.Get(DatabaseDef).(*pgxpool.Pool)
			ctx, _ := ctn.Get(ContextDef).(context.Context)

			conn, err := pool.Acquire(ctx)
			if err != nil {
				return nil, err
			}

			return users.NewRepo(conn), nil
		},
	}
}
