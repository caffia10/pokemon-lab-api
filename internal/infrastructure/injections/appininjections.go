package injections

import (
	pkmApp "pokemon-lab-api/internal/application"
	pkmtApp "pokemon-lab-api/internal/application"
)

func RetrieveAppInjections() []interface{} {
	return []interface{}{
		pkmApp.NewCreatePokemonInBulkUsecase,
		pkmApp.NewCreatePokemonUsecase,
		pkmApp.NewRetrievePokemonbyIdUsecase,
		pkmtApp.NewRetrieveAllPokemonTypeUsecase,
		pkmtApp.NewCreatePokemonTypeInBulkUsecase,
		pkmtApp.NewCreatePokemonTypeUsecase,
	}
}
