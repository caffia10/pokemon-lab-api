package application

import (
	"pokemon-lab-api/internal/pokemon/domain"
	"pokemon-lab-api/pkg/mderrors"

	"go.uber.org/zap"
)

// defaultRetrieveAllPokemonCase implements RetrievePokemonbyIdUseCase
type defaultRetrieveAllPokemonCase struct {
	repo   domain.PokemonRepository
	logger *zap.Logger
}

func (uc *defaultRetrieveAllPokemonCase) Do() ([]Pokemon, error) {

	p, errU := uc.repo.RetrieveAll()

	if errU != nil {
		return nil, mderrors.NewMetadataError(errU)
	}

	return p, nil
}

func NewRetrieveAllPokemonTypeUsecase(r domain.PokemonRepository, l *zap.Logger) RetrieveAllPokemonUseCase {

	return &defaultRetrieveAllPokemonCase{
		repo:   r,
		logger: l,
	}
}
