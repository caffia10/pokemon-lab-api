package controller

import (
	"pokemon-lab-api/internal/application"
	"pokemon-lab-api/internal/domain"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type CreatePokemon struct {
	s      application.CreatePokemonUseCase
	logger *zap.Logger
}

func (c *CreatePokemon) Validate(ctx *fiber.Ctx) error {

	pkm := new(pokemonDto)
	errB := ctx.BodyParser(pkm)

	if errB != nil {
		lf := []zap.Field{
			zap.String("logger", "CreatePokemonController"),
			zap.String("sub-logger", "Validate"),
		}
		lf = append(lf, zap.NamedError("error", errB))
		c.logger.Error("unexpected error at unmarshaling request", lf...)
		return fiber.ErrBadRequest
	}

	ctx.Locals("request", pkm)
	return ctx.Next()
}

func (c *CreatePokemon) Handle(ctx *fiber.Ctx) error {

	pkmd, ok := ctx.Locals("request").(*pokemonDto)

	lf := []zap.Field{
		zap.String("logger", "CreatePokemonController"),
		zap.String("sub-logger", "Handle"),
	}

	if !ok {
		c.logger.Error("error at retrieving request from locals", lf...)
		return fiber.ErrBadRequest
	}

	pkm := &domain.Pokemon{
		Name:     pkmd.Name,
		Category: pkmd.Category,
		Weight:   pkmd.Weight,
		ImgUrl:   pkmd.ImgUrl,
		Types:    make([]domain.PokemonType, len(pkmd.Types)),
	}

	if len(pkmd.Evolutions) > 0 {

		pkm.Evolutions = make([]domain.Evolution, len(pkm.Evolutions))
		for i, ev := range pkmd.Evolutions {

			pkm.Evolutions[i] = domain.Evolution{
				Id:          ev.Id,
				Requirement: ev.Requirement,
			}
		}
	}

	for i, t := range pkmd.Types {
		pkm.Types[i] = domain.PokemonType(t)
	}

	errC := c.s.Do(pkm)

	if errC != nil {
		lf = append(lf, zap.NamedError("error", errC))
		c.logger.Error("error at creating pokemon request from service", lf...)
		return fiber.ErrConflict
	}

	return ctx.SendStatus(fiber.StatusOK)
}

func NewCreatePokemonController(s application.CreatePokemonUseCase, logger *zap.Logger) *CreatePokemon {
	return &CreatePokemon{
		s:      s,
		logger: logger,
	}
}
