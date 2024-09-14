package injections

import (
	pkmDomainUc "pokemon-lab-api/internal/pokemon/domain/usecase"
)

func RetrieveDomainInjections() []interface{} {
	return []interface{}{
		pkmDomainUc.NewRetrievePokemonbyIdUsecase,
		pkmDomainUc.NewCreatePokemonInBulkUsecase,
		pkmDomainUc.NewCreatePokemonUsecase,
	}
}
