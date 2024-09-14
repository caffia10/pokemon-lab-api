package usecase

import (
	"pokemon-lab-api/internal/pokemon/domain"

	"go.uber.org/zap"
)

type createPokemonInBulk struct {
	repo   domain.PokemonRepository
	logger *zap.Logger
}

func (uc *createPokemonInBulk) Do(pkms []*domain.Pokemon) error {

	err := uc.repo.CreateManyPokemon(pkms)

	if err != nil {

		lf := []zap.Field{
			zap.String("logger", "createPokemonInBulk"),
			zap.String("sub-logger", "Do"),
			zap.NamedError("error", err),
		}
		uc.logger.Error("error at creating pokemon", lf...)
	}

	return err
}

func NewCreatePokemonInBulkUsecase(r domain.PokemonRepository, logger *zap.Logger) domain.CreatePokemonInBulkUsecase {

	return &createPokemonInBulk{
		repo:   r,
		logger: logger,
	}
}
