package definitions

import (
	"context"

	"github.com/sarulabs/di"
)

const (
	ContextDef = "context"
)

func getContextDef() di.Def {
	return di.Def{
		Name:  ContextDef,
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return context.Background(), nil
		},
	}
}
