package pokemon

import (
	"go.uber.org/zap"
)

type Service struct {
	repo   PokemonRepository
	logger *zap.Logger
}

func (uc *Service) Create(pkm *Pokemon) error {

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

func (uc *Service) CreateBulk(pkms []*Pokemon) error {

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

func (uc *Service) RetriveById(id string) (*Pokemon, error) {

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

	return p, nil
}

func NewCreatePokemonUsecase(r PokemonRepository, logger *zap.Logger) *Service {

	return &Service{
		repo:   r,
		logger: logger,
	}
}
