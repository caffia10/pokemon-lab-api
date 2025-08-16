package http

import (
	"pokemon-lab-api/internal/pokemon/application"
	"pokemon-lab-api/pkg/mderrors"

	"github.com/gofiber/fiber/v2"
)

type RetrievePokemonController struct {
	s application.RetrievePokemonbyIdUseCase
}

func (c *RetrievePokemonController) Handle(ctx *fiber.Ctx) error {

	id := ctx.Params("id")

	pkmts, errR := c.s.Do(id)

	if errR != nil {
		return mderrors.NewMetadataError(errR)
	}

	return ctx.JSON(pkmts)
}

func NewRetrievePokemonController(s application.RetrievePokemonbyIdUseCase) *RetrievePokemonController {
	return &RetrievePokemonController{
		s: s,
	}
}
