package bootstrap

import (
	"pokemon-lab-api/internal/config"
	"pokemon-lab-api/internal/http"
	"pokemon-lab-api/internal/logger"
	"pokemon-lab-api/internal/scylladb"
)

func RetrieveInjections() []interface{} {
	return []interface{}{
		scylladb.NewPokemonRepository,
		http.NewFiberApiServer,
		config.NewConfig,
		scylladb.NewPokemonTypeRepository,
		scylladb.NewSession,
		logger.New,
	}
}
