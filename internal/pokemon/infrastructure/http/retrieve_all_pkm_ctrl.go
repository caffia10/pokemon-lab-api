package http

import (
	"pokemon-lab-api/internal/pokemon/application"
	"pokemon-lab-api/pkg/mderrors"

	"github.com/gofiber/fiber/v2"
)

type RetrieveAllPokemonController struct {
	s application.RetrieveAllPokemonUseCase
}

func (c *RetrieveAllPokemonController) Handle(ctx *fiber.Ctx) error {

	pkmts, errR := c.s.Do()

	if errR != nil {
		return mderrors.NewMetadataError(errR)
	}

	return ctx.JSON(pkmts)
}

func NewRetrieveAllPokemonController(s application.RetrieveAllPokemonUseCase) *RetrieveAllPokemonController {
	return &RetrieveAllPokemonController{
		s: s,
	}
}
