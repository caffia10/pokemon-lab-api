package application

import (
	pkmtDomain "pokemon-lab-api/internal/pokemon-type/domain"
	"pokemon-lab-api/internal/pokemon/domain"
)

type Pokemon = domain.Pokemon
type PokemonType = pkmtDomain.PokemonType

type RetrievePokemonbyIdUseCase interface {
	Do(id string) (*Pokemon, error)
}

type CreatePokemonUseCase interface {
	Do(pkm *Pokemon) error
}

type CreatePokemonInBulkUseCase interface {
	Do(pkms []*Pokemon) error
}
