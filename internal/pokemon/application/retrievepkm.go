package application

import (
	pkmtDomain "pokemon-lab-api/internal/pokemon-type/domain"
	"pokemon-lab-api/internal/pokemon/domain"

	"go.uber.org/zap"
)

// defaultRetrievePokemonbyIdUseCase implements RetrievePokemonbyIdUseCase
type defaultRetrievePokemonbyIdUseCase struct {
	repo     domain.PokemonRepository
	pkmtRepo pkmtDomain.PokemonTypeRepository
	logger   *zap.Logger
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

	{
		tIds := make([]string, len(p.Types))
		for _, t := range p.Types {
			tIds = append(tIds, t.Id)
		}

		ts, errRts := uc.pkmtRepo.RetriveByIds(tIds)

		if errRts != nil {
			lf = append(lf, zap.NamedError("error", errRts))
			uc.logger.Error("error at retriving pokemon types", lf...)
			return nil, errRts
		}

		copy(p.Types, ts)
	}

	return p, nil
}

func NewRetrievePokemonbyIdUsecase(r domain.PokemonRepository, pkmtr pkmtDomain.PokemonTypeRepository) RetrievePokemonbyIdUseCase {

	return &defaultRetrievePokemonbyIdUseCase{
		repo:     r,
		pkmtRepo: pkmtr,
	}
}
