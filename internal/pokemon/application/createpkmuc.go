package application

import (
	"pokemon-lab-api/internal/pokemon/domain"

	"go.uber.org/zap"
)

// defaultCreatePokemonUseCase implements CreatePokemonUseCase
type defaultCreatePokemonUseCase struct {
	repo   domain.PokemonRepository
	logger *zap.Logger
}

func (uc *defaultCreatePokemonUseCase) Do(pkm *domain.Pokemon) error {

	lf := []zap.Field{
		zap.String("logger", "createPokemon"),
		zap.String("sub-logger", "Do"),
		zap.String("pokemon-name", pkm.Name),
	}

	uc.logger.Info("creating pokemon", lf...)

	err := uc.repo.Create(pkm)

	if err != nil {
		lf = append(lf, zap.NamedError("error", err))
		uc.logger.Error("error at creating the pokemon", lf...)
	}

	return err
}

func NewCreatePokemonUsecase(r domain.PokemonRepository, logger *zap.Logger) CreatePokemonUseCase {

	return &defaultCreatePokemonUseCase{
		repo:   r,
		logger: logger,
	}
}
