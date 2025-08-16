package application

import (
	"pokemon-lab-api/internal/pokemon-type/domain"
	"pokemon-lab-api/pkg/mderrors"

	"go.uber.org/zap"
)

// defaultRetrieveAllPokemonTypeCase implements RetrievePokemonbyIdUseCase
type defaultRetrieveAllPokemonTypeCase struct {
	repo   domain.PokemonTypeRepository
	logger *zap.Logger
}

func (uc *defaultRetrieveAllPokemonTypeCase) Do() ([]*PokemonType, error) {

	p, errU := uc.repo.RetriveAll()

	if errU != nil {
		return nil, mderrors.NewMetadataError(errU)
	}

	return p, nil
}

func NewRetrieveAllPokemonTypeUsecase(r domain.PokemonTypeRepository, l *zap.Logger) RetrieveAllPokemonTypeUseCase {

	return &defaultRetrieveAllPokemonTypeCase{
		repo:   r,
		logger: l,
	}
}
