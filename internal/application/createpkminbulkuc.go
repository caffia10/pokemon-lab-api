package application

import (
	"pokemon-lab-api/internal/domain"

	"go.uber.org/zap"
)

// defaultCreatePokemonInBulkUseCase implements CreatePokemonInBulkUseCase
type defaultCreatePokemonInBulkUseCase struct {
	repo   domain.PokemonRepository
	logger *zap.Logger
}

func (uc *defaultCreatePokemonInBulkUseCase) Do(pkms []domain.Pokemon) error {

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

func NewCreatePokemonInBulkUsecase(r domain.PokemonRepository, logger *zap.Logger) CreatePokemonInBulkUseCase {

	return &defaultCreatePokemonInBulkUseCase{
		repo:   r,
		logger: logger,
	}
}
