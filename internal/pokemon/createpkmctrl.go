package pokemon

import (
	pokemontype "pokemon-lab-api/internal/pokemon-type"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type CreatePokemonController struct {
	s      *Service
	logger *zap.Logger
}

func (c *CreatePokemonController) Validate(ctx *fiber.Ctx) error {

	pkm := new(pokemonRequestDto)
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

	pkmd, ok := ctx.Locals("request").(*pokemonRequestDto)

	lf := []zap.Field{
		zap.String("logger", "CreatePokemonController"),
		zap.String("sub-logger", "Handle"),
	}

	if !ok {
		c.logger.Error("error at retrieving request from locals", lf...)
		return fiber.ErrBadRequest
	}

	pkm := &Pokemon{
		Name:     pkmd.Name,
		Category: pkmd.Category,
		Weight:   pkmd.Weight,
		ImgUrl:   pkmd.ImgUrl,
		Types:    make([]*pokemontype.PokemonType, 0, len(pkmd.Types)),
	}

	for _, v := range pkmd.Types {
		pkm.Types = append(pkm.Types, (*pokemontype.PokemonType)(v))
	}

	errC := c.s.Create(pkm)

	if errC != nil {
		lf = append(lf, zap.NamedError("error", errC))
		c.logger.Error("error at creating pokemon request from service", lf...)
		return fiber.ErrConflict
	}

	return ctx.SendStatus(fiber.StatusOK)
}

func NewCreatePokemonController(s *Service, logger *zap.Logger) *CreatePokemonController {
	return &CreatePokemonController{
		s:      s,
		logger: logger,
	}
}
