package application

import (
	"pokemon-lab-api/internal/pokemon-type/domain"

	"go.uber.org/zap"
)

// defaultRetrieveAllPokemonTypeCase implements RetrievePokemonbyIdUseCase
type defaultRetrieveAllPokemonTypeCase struct {
	repo   domain.PokemonTypeRepository
	logger *zap.Logger
}

func (uc *defaultRetrieveAllPokemonTypeCase) Do() ([]*PokemonType, error) {

	lf := []zap.Field{
		zap.String("logger", "defaultRetrieveAllPokemonTypeCase"),
		zap.String("sub-logger", "Do"),
	}

	uc.logger.Info("retrieving pokemon type", lf...)

	p, errU := uc.repo.RetriveAll()

	if errU != nil {
		lf = append(lf, zap.NamedError("error", errU))
		uc.logger.Error("error at retriving pokemon type", lf...)
		return nil, errU
	}

	return p, nil
}

func NewRetrieveAllPokemonTypeUsecase(r domain.PokemonTypeRepository, l *zap.Logger) RetrieveAllPokemonTypeUseCase {

	return &defaultRetrieveAllPokemonTypeCase{
		repo:   r,
		logger: l,
	}
}
