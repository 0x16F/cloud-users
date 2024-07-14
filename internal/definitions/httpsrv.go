package definitions

import (
	"github.com/0x16F/cloud/users/internal/controller/httpsrv"
	"github.com/0x16F/cloud/users/internal/controller/httpsrv/handlers/users"
	"github.com/sarulabs/di"
)

const (
	HTTPServerDef = "http-server"
)

func getHTTPServerDef() di.Def {
	return di.Def{
		Name:  HTTPServerDef,
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			server := httpsrv.NewServer()

			api := server.App.Group("/api/v1")

			usersHandler, _ := ctn.Get(UsersHandlerDef).(*users.Handler)
			usersHandler.RegisterRoutes(api.Group("/users"))

			return server, nil
		},
		Close: func(obj interface{}) error {
			server, _ := obj.(*httpsrv.Server)
			return server.Stop()
		},
	}
}
