package definitions

import (
	"github.com/sarulabs/di"
)

func New() (di.Container, error) {
	builder, err := di.NewBuilder()
	if err != nil {
		return nil, err
	}

	if err := builder.Add([]di.Def{
		getContextDef(),
		getConfigDef(),
		getLoggerDef(),
		getErrorsServiceDef(),
		getDatabaseDef(),
		getUsersRepoDef(),
		getUsersServiceDef(),
		getUsersHandlerDef(),
		getHTTPServerDef(),
	}...); err != nil {
		return nil, err
	}

	return builder.Build(), nil
}
