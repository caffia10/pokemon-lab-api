package pokemontype

import (
	"sync"

	"go.uber.org/zap"
)

type Service struct {
	repo   PokemonTypeRepository
	logger *zap.Logger
}

func (uc *Service) Create(pkm *PokemonType) error {

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

func (uc *Service) CreateBulk(pkmst []*PokemonType) error {

	var wg sync.WaitGroup
	errChannel := make(chan *PokemonType)
	for _, pkmt := range pkmst {
		wg.Add(1)
		go func(pt *PokemonType) {
			defer wg.Done()
			err := uc.repo.Create(pt)

			if err != nil {
				lf := []zap.Field{
					zap.String("logger", "createPokemonInBulk"),
					zap.String("sub-logger", "Do"),
					zap.NamedError("error", err),
				}
				uc.logger.Error("error at creating pokemon", lf...)
				errChannel <- pt
			}
		}(pkmt)
	}

	wg.Wait()
	close(errChannel)

	if len(errChannel) > 0 {

		if len(errChannel) == len(pkmst) {
			return AllPokemonTypesFail
		}

		ept := make([]*PokemonType, 0)
		for pt := range errChannel {

			ept = append(ept, pt)
		}

		return NewSomePokemonFailCreationError(ept)
	}

	return nil
}

func (uc *Service) RetrieveAll() ([]*PokemonType, error) {

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

func NewCreatePokemonTypeUsecase(r PokemonTypeRepository, logger *zap.Logger) *Service {

	return &Service{
		repo:   r,
		logger: logger,
	}
}
