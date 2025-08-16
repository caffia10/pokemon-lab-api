package injections

import (
	pkmtApp "pokemon-lab-api/internal/pokemon-type/application"
	pkmApp "pokemon-lab-api/internal/pokemon/application"
)

func RetrieveAppInjections() []interface{} {
	return []interface{}{
		pkmApp.NewCreatePokemonInBulkUsecase,
		pkmApp.NewCreatePokemonUsecase,
		pkmApp.NewRetrievePokemonbyIdUsecase,
		pkmApp.NewRetrieveAllPokemonTypeUsecase,
		pkmtApp.NewRetrieveAllPokemonTypeUsecase,
		pkmtApp.NewCreatePokemonTypeInBulkUsecase,
		pkmtApp.NewCreatePokemonTypeUsecase,
	}
}
