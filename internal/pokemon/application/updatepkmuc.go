package application

import (
	"pokemon-lab-api/internal/pokemon/domain"

	"go.uber.org/zap"
)

// defaultUpdatePokemonUseCase implements UpdatePokemonUseCase
type defaultUpdatePokemonUseCase struct {
	repo   domain.PokemonRepository
	logger *zap.Logger
}

func (uc *defaultUpdatePokemonUseCase) Do(pkm *domain.Pokemon) error {

	lf := []zap.Field{
		zap.String("logger", "updatePokemon"),
		zap.String("sub-logger", "Do"),
		zap.String("pokemon-name", pkm.Name),
	}

	uc.logger.Info("updating pokemon", lf...)

	err := uc.repo.Update(pkm)

	if err != nil {
		lf = append(lf, zap.NamedError("error", err))
		uc.logger.Error("error at creating the pokemon", lf...)
	}

	return err
}

func NewUpdatePokemonUsecase(r domain.PokemonRepository, logger *zap.Logger) UpdatePokemonUseCase {

	return &defaultUpdatePokemonUseCase{
		repo:   r,
		logger: logger,
	}
}
