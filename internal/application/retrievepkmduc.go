package application

import (
	"pokemon-lab-api/internal/domain"
	pkmtDomain "pokemon-lab-api/internal/domain"

	"go.uber.org/zap"
)

// defaultRetrievePokemonDetailedUseCase implements RetrievePokemonbyIdUseCase
type defaultRetrievePokemonDetailedUseCase struct {
	repo   domain.PokemonRepository
	logger *zap.Logger
}

func (uc *defaultRetrievePokemonDetailedUseCase) Do(id string) (*domain.Pokemon, error) {

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

	uc.LoadEvolution(p)

	return p, nil
}

func (uc *defaultRetrievePokemonDetailedUseCase) LoadEvolution(pkm *domain.Pokemon) error {
	if len(pkm.Evolutions) == 0 {
		return nil
	}

	lf := []zap.Field{
		zap.String("logger", "defaultRetrievePokemonDetailedUseCase"),
		zap.String("sub-logger", "RetrieveEvolution"),
	}

	for _, ev := range pkm.Evolutions {
		p, errU := uc.repo.RetriveById(ev.Id)
		if errU != nil {
			lf = append(lf, zap.String("pokemon-id", p.Id), zap.NamedError("error", errU))
			uc.logger.Error("error at retriving pokemon", lf...)
			return errU
		}

		// Maybe could be a sub load by go rutine

		if errLoadEvolution := uc.LoadEvolution(p); errLoadEvolution != nil {
			lf = append(lf, zap.String("pokemon-id", p.Id), zap.NamedError("error", errU))
			uc.logger.Error("error at loading pokemon nested evolution", lf...)
			return errU
		}

		ev.Pokemon = p

	}

	return nil
}

func NewRetrievePokemonDetailedUseCase(r domain.PokemonRepository, pkmtr pkmtDomain.PokemonTypeRepository) RetrievePokemonbyIdUseCase {

	return &defaultRetrievePokemonDetailedUseCase{
		repo: r,
	}
}
