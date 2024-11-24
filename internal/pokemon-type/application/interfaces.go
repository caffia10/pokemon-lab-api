package application

import (
	pkmtDomain "pokemon-lab-api/internal/pokemon-type/domain"
)

type PokemonType = pkmtDomain.PokemonType

type RetrieveAllPokemonTypeUseCase interface {
	Do() ([]*PokemonType, error)
}

type CreatePokemonTypeUseCase interface {
	Do(pkm *PokemonType) error
}

type CreatePokemonTypeInBulkUseCase interface {
	Do(pkms []*PokemonType) error
}
