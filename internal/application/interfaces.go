package application

import (
	"pokemon-lab-api/internal/domain"
)

type RetrievePokemonbyIdUseCase interface {
	Do(id string) (*domain.Pokemon, error)
}

type CreatePokemonUseCase interface {
	Do(pkm *domain.Pokemon) error
}

type CreatePokemonInBulkUseCase interface {
	Do(pkms []domain.Pokemon) error
}

type RetrieveAllPokemonTypeUseCase interface {
	Do() ([]domain.PokemonType, error)
}

type CreatePokemonTypeUseCase interface {
	Do(pkm *domain.PokemonType) error
}

type CreatePokemonTypeInBulkUseCase interface {
	Do(pkms []domain.PokemonType) error
}
