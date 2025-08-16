package application

import (
	"pokemon-lab-api/internal/pokemon/domain"
	"pokemon-lab-api/pkg/mderrors"
)

// defaultCreatePokemonInBulkUseCase implements CreatePokemonInBulkUseCase
type defaultCreatePokemonInBulkUseCase struct {
	repo domain.PokemonRepository
}

func (uc *defaultCreatePokemonInBulkUseCase) Do(pkms []domain.Pokemon) error {

	err := uc.repo.CreateManyPokemon(pkms)

	if err != nil {
		return mderrors.NewMetadataError(err)
	}

	return nil
}

func NewCreatePokemonInBulkUsecase(r domain.PokemonRepository) CreatePokemonInBulkUseCase {

	return &defaultCreatePokemonInBulkUseCase{
		repo: r,
	}
}
