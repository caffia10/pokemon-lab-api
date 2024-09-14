package usecase

import (
	"pokemon-lab-api/internal/pokemon/domain"

	"go.uber.org/zap"
)

type createPokemon struct {
	repo   domain.PokemonRepository
	logger *zap.Logger
}

func (uc *createPokemon) Do(pkm *domain.Pokemon) error {

	lf := []zap.Field{
		zap.String("logger", "createPokemon"),
		zap.String("sub-logger", "Do"),
		zap.String("pokemon-id", pkm.Id),
	}

	uc.logger.Info("creating pokemon", lf...)

	err := uc.repo.CreatePokemon(pkm)

	if err != nil {
		lf = append(lf, zap.NamedError("error", err))
		uc.logger.Error("error at creating the pokemon", lf...)
	}

	return err
}

func NewCreatePokemonUsecase(r domain.PokemonRepository, logger *zap.Logger) domain.CreatePokemonUsecase {

	return &createPokemon{
		repo:   r,
		logger: logger,
	}
}
