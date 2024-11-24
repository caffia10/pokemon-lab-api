package application

import (
	"pokemon-lab-api/internal/pokemon-type/domain"
	"sync"

	"go.uber.org/zap"
)

// defaultCreatePokemonTypeInBulkUseCase implements CreatePokemonInBulkUseCase
type defaultCreatePokemonTypeInBulkUseCase struct {
	repo   domain.PokemonTypeRepository
	logger *zap.Logger
}

func (uc *defaultCreatePokemonTypeInBulkUseCase) Do(pkmst []*PokemonType) error {

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

func NewCreatePokemonTypeInBulkUsecase(r domain.PokemonTypeRepository, logger *zap.Logger) CreatePokemonTypeInBulkUseCase {

	return &defaultCreatePokemonTypeInBulkUseCase{
		repo:   r,
		logger: logger,
	}
}
