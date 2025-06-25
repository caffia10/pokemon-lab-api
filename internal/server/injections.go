package server

import (
	"pokemon-lab-api/internal/config"
	"pokemon-lab-api/internal/pokemon"
	pokemontype "pokemon-lab-api/internal/pokemon-type"
)

func RetrieveInfraInjections() []interface{} {
	return []interface{}{
		pokemon.NewPokemonRepository,
		NewFiberApiServer,
		NewMongoDatabase,
		config.NewConfig,
		pokemontype.NewPokemonTypeRepository,
		pokemon.NewCreatePokemonController,
		NewSession,
		New,
		pokemontype.NewCreatePokemonTypeController,
		pokemontype.NewRetrieveAllPokemonTypeController,
		pokemontype.NewCreatePokemonTypeUsecase,
		pokemon.NewCreatePokemonUsecase,
	}
}
