package http

import (
	"pokemon-lab-api/internal/pokemon-type/application"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type CreatePokemonTypeController struct {
	s      application.CreatePokemonTypeUseCase
	logger *zap.Logger
}

func (c *CreatePokemonTypeController) Validate(ctx *fiber.Ctx) error {

	pkmt := new(pokemonTypeDto)
	errB := ctx.BodyParser(pkmt)

	if errB != nil {
		lf := []zap.Field{
			zap.String("logger", "CreatePokemonTypeController"),
			zap.String("sub-logger", "Validate"),
		}
		lf = append(lf, zap.NamedError("error", errB))
		c.logger.Error("unexpected error at unmarshaling request", lf...)
		return fiber.ErrBadRequest
	}

	ctx.Locals("request", pkmt)
	return ctx.Next()
}

func (c *CreatePokemonTypeController) Handle(ctx *fiber.Ctx) error {

	pkmd, ok := ctx.Locals("request").(*pokemonTypeDto)

	lf := []zap.Field{
		zap.String("logger", "CreatePokemonController"),
		zap.String("sub-logger", "Handle"),
	}

	if !ok {
		c.logger.Error("error at retrieving request from locals", lf...)
		return fiber.ErrBadRequest
	}

	errC := c.s.Do((*application.PokemonType)(pkmd))

	if errC != nil {
		lf = append(lf, zap.NamedError("error", errC))
		c.logger.Error("error at creating pokemon request from service", lf...)
		return fiber.ErrConflict
	}

	return ctx.SendStatus(fiber.StatusOK)
}

func NewCreatePokemonTypeController(s application.CreatePokemonTypeUseCase, logger *zap.Logger) *CreatePokemonTypeController {
	return &CreatePokemonTypeController{
		s:      s,
		logger: logger,
	}
}
