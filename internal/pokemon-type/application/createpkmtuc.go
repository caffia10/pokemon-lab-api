package application

import (
	"pokemon-lab-api/internal/pokemon-type/domain"

	"go.uber.org/zap"
)

// defaultCreatePokemonUseCase implements CreatePokemonUseCase
type defaultCreatePokemonTypeUseCase struct {
	repo   domain.PokemonTypeRepository
	logger *zap.Logger
}

func (uc *defaultCreatePokemonTypeUseCase) Do(pkm *PokemonType) error {

	lf := []zap.Field{
		zap.String("logger", "createPokemon"),
		zap.String("sub-logger", "Do"),
		zap.String("pokemon-type-name", pkm.Name),
	}

	uc.logger.Info("creating pokemon", lf...)

	err := uc.repo.Create(pkm)

	if err != nil {
		lf = append(lf, zap.NamedError("error", err))
		uc.logger.Error("error at creating the pokemon type", lf...)
	}

	return err
}

func NewCreatePokemonTypeUsecase(r domain.PokemonTypeRepository, logger *zap.Logger) CreatePokemonTypeUseCase {

	return &defaultCreatePokemonTypeUseCase{
		repo:   r,
		logger: logger,
	}
}
