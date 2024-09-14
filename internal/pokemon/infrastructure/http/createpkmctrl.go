package http

import (
	"pokemon-lab-api/internal/pokemon/application"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type CreatePokemonController struct {
	s      application.CreatePokemonService
	logger *zap.Logger
}

func (c *CreatePokemonController) Validate(ctx *fiber.Ctx) error {

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

func (c *CreatePokemonController) Handle(ctx *fiber.Ctx) error {

	pkmd, ok := ctx.Locals("request").(*pokemonDto)

	lf := []zap.Field{
		zap.String("logger", "CreatePokemonController"),
		zap.String("sub-logger", "Handle"),
	}

	if !ok {
		c.logger.Error("error at retrieving request from locals", lf...)
		return fiber.ErrBadRequest
	}

	pkm := &application.Pokemon{
		Name:     pkmd.Name,
		Category: pkmd.Category,
		Weight:   pkmd.Weight,
		ImgUrl:   pkmd.ImgUrl,
		Types:    make([]*application.PokemonType, len(pkmd.Types)),
	}

	for _, v := range pkmd.Types {
		pkm.Types = append(pkm.Types, (*application.PokemonType)(v))
	}

	errC := c.s.Do(pkm)

	if errC != nil {
		lf = append(lf, zap.NamedError("error", errC))
		c.logger.Error("error at creating pokemon request from service", lf...)
		return fiber.ErrConflict
	}

	return ctx.SendStatus(fiber.StatusOK)
}

func NewCreatePokemonController(s application.CreatePokemonService, logger *zap.Logger) *CreatePokemonController {
	return &CreatePokemonController{
		s:      s,
		logger: logger,
	}
}
