package application

import (
	pkmtDomain "pokemon-lab-api/internal/pokemon-type/domain"
	"pokemon-lab-api/internal/pokemon/domain"
	"pokemon-lab-api/internal/pokemon/domain/usecase"
)

type Pokemon = domain.Pokemon
type PokemonType = domain.PokemonType

type RetrievePokemonbyIdService interface {
	Do(id string) (*Pokemon, error)
}

func NewRetrievePokemonbyIdService(r domain.PokemonRepository, pkmtr pkmtDomain.PokemonTypeRepository) RetrievePokemonbyIdService {
	return usecase.NewRetrievePokemonbyIdUsecase(r, pkmtr)
}

type CreatePokemonService interface {
	Do(pkm *Pokemon) error
}

func NewCreatePokemonService(r domain.PokemonRepository) CreatePokemonService {
	return usecase.NewCreatePokemonUsecase(r)
}

type CreatePokemonInBulkService interface {
	Do(pkms []*Pokemon) error
}
