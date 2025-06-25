package injections

import (
	"pokemon-lab-api/internal/infrastructure/config"
	"pokemon-lab-api/internal/infrastructure/controller"
	"pokemon-lab-api/internal/infrastructure/fiberapi"
	"pokemon-lab-api/internal/infrastructure/logger"
	"pokemon-lab-api/internal/infrastructure/mongoctx"
	"pokemon-lab-api/internal/infrastructure/repository"
	"pokemon-lab-api/internal/infrastructure/scylladbctx"
)

func RetrieveInfraInjections() []interface{} {
	return []interface{}{
		repository.NewPokemonRepository,
		fiberapi.NewFiberApiServer,
		mongoctx.NewMongoDatabase,
		config.NewConfig,
		repository.NewPokemonTypeRepository,
		controller.NewCreatePokemonController,
		scylladbctx.NewSession,
		logger.New,
		controller.NewCreatePokemonTypeController,
		controller.NewRetrieveAllPokemonTypeController,
	}
}
