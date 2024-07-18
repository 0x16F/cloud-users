package definitions

import (
	"github.com/sarulabs/di/v2"
)

func New() (di.Container, error) {
	builder, err := di.NewEnhancedBuilder()
	if err != nil {
		return di.Container{}, err
	}

	if err := addDef(builder,
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
	); err != nil {
		return di.Container{}, err
	}

	return builder.Build()
}

func addDef(builder *di.EnhancedBuilder, defs ...*di.Def) error {
	for _, def := range defs {
		if err := builder.Add(def); err != nil {
			return err
		}
	}

	return nil
}
