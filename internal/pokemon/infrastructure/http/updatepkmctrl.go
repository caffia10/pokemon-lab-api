package http

import (
	"pokemon-lab-api/internal/pokemon/application"
	"pokemon-lab-api/pkg/mderrors"

	"github.com/gofiber/fiber/v2"
)

type UpdatePokemonController struct {
	s application.UpdatePokemonUseCase
}

func (c *UpdatePokemonController) Validate(ctx *fiber.Ctx) error {

	pkm := new(pokemonDto)
	errB := ctx.BodyParser(pkm)

	if errB != nil {
		return fiber.ErrBadRequest
	}

	ctx.Locals("request", pkm)
	return ctx.Next()
}

func (c *UpdatePokemonController) Handle(ctx *fiber.Ctx) error {

	pkmd, ok := ctx.Locals("request").(*pokemonDto)
	if !ok {
		return fiber.ErrBadRequest
	}

	pkm := &application.Pokemon{
		Name:     pkmd.Name,
		Category: pkmd.Category,
		Weight:   pkmd.Weight,
		ImgUrl:   pkmd.ImgUrl,
		Types:    make([]application.PokemonType, 0, len(pkmd.Types)),
	}

	for _, v := range pkmd.Types {
		pkm.Types = append(pkm.Types, (application.PokemonType)(v))
	}

	errC := c.s.Do(pkm)

	if errC != nil {
		return mderrors.NewMetadataError(errC)
	}

	return ctx.SendStatus(fiber.StatusOK)
}

func NewUpdatePokemonController(s application.UpdatePokemonUseCase) *UpdatePokemonController {
	return &UpdatePokemonController{
		s: s,
	}
}
