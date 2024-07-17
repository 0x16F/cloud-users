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

		getFFlagsProviderDef(),
		getFFlagsClientDef(),

		getDatabaseDef(),
		getUsersRepoDef(),

		getErrorsServiceDef(),
		getUsersServiceDef(),
		getFFlagsServiceDef(),

		getHTTPServerDef(),
		getUsersHandlerDef(),
		getFeaturesServiceDef(),
	}...); err != nil {
		return nil, err
	}

	return builder.Build(), nil
}
