package definitions

import (
	"github.com/0x16F/cloud-users/internal/controller/httpsrv"
	"github.com/0x16F/cloud-users/internal/controller/httpsrv/handlers/users"
	"github.com/sarulabs/di"
)

const (
	HTTPServerDef = "http_server"
)

func getHTTPServerDef() di.Def {
	return di.Def{
		Name:  HTTPServerDef,
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			usersHandler, _ := ctn.Get(UsersHandlerDef).(*users.Handler)

			server := httpsrv.NewServer()

			v1 := server.App.Group("/api/v1")
			{
				users := v1.Group("/users")
				{
					users.Get("/", usersHandler.GetUsers)
					users.Get("/:id", usersHandler.GetUser)
					users.Post("/", usersHandler.CreateUser)
					users.Patch("/:id/email", usersHandler.UpdateEmail)
					users.Patch("/:id/username", usersHandler.UpdateUsername)
					users.Patch("/:id/password", usersHandler.UpdatePassword)
					users.Delete("/:id", usersHandler.DeleteUser)
				}
			}

			return server, nil
		},
		Close: func(obj interface{}) error {
			server, _ := obj.(*httpsrv.Server)
			return server.Stop()
		},
	}
}
