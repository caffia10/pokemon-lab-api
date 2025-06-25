package application

import (
	"pokemon-lab-api/internal/domain"
	pkmtDomain "pokemon-lab-api/internal/domain"

	"go.uber.org/zap"
)

// defaultRetrievePokemonbyIdUseCase implements RetrievePokemonbyIdUseCase
type defaultRetrievePokemonbyIdUseCase struct {
	repo   domain.PokemonRepository
	logger *zap.Logger
}

func (uc *defaultRetrievePokemonbyIdUseCase) Do(id string) (*domain.Pokemon, error) {

	lf := []zap.Field{
		zap.String("logger", "retrievePokemonbyId"),
		zap.String("sub-logger", "Do"),
		zap.String("pokemon-id", id),
	}

	uc.logger.Info("retrieving pokemon", lf...)

	p, errU := uc.repo.RetriveById(id)

	if errU != nil {
		lf = append(lf, zap.NamedError("error", errU))
		uc.logger.Error("error at retriving pokemon", lf...)
		return nil, errU
	}

	return p, nil
}

func NewRetrievePokemonbyIdUsecase(r domain.PokemonRepository, pkmtr pkmtDomain.PokemonTypeRepository) RetrievePokemonbyIdUseCase {

	return &defaultRetrievePokemonbyIdUseCase{
		repo: r,
	}
}
