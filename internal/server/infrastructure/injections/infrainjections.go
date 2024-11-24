package injections

import (
	pkmtInfraHttp "pokemon-lab-api/internal/pokemon-type/infrastructure/http"
	pkmtInfraRepo "pokemon-lab-api/internal/pokemon-type/infrastructure/repo/scyllarepo"
	pkmInfraHttp "pokemon-lab-api/internal/pokemon/infrastructure/http"
	pkmInfraRepo "pokemon-lab-api/internal/pokemon/infrastructure/repo/scyllarepo"
	serverInfra "pokemon-lab-api/internal/server/infrastructure"
	"pokemon-lab-api/internal/server/infrastructure/config"
	"pokemon-lab-api/internal/server/infrastructure/logger"
	"pokemon-lab-api/internal/server/infrastructure/mongoctx"
	"pokemon-lab-api/internal/server/infrastructure/scylladbctx"
)

func RetrieveInfraInjections() []interface{} {
	return []interface{}{
		pkmInfraRepo.NewPokemonRepository,
		serverInfra.NewFiberApiServer,
		mongoctx.NewMongoDatabase,
		config.NewConfig,
		pkmtInfraRepo.NewPokemonTypeRepository,
		pkmInfraHttp.NewCreatePokemonController,
		scylladbctx.NewSession,
		logger.New,
		pkmtInfraHttp.NewCreatePokemonTypeController,
		pkmtInfraHttp.NewRetrieveAllPokemonTypeController,
	}
}
