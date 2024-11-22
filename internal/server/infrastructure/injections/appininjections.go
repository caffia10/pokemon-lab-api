package injections

import (
	pkmApp "pokemon-lab-api/internal/pokemon/application"
)

func RetrieveAppInjections() []interface{} {
	return []interface{}{
		pkmApp.NewCreatePokemonInBulkUsecase,
		pkmApp.NewCreatePokemonUsecase,
		pkmApp.NewRetrievePokemonbyIdUsecase,
	}
}
