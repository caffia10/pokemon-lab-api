package http

import (
	"pokemon-lab-api/internal/pokemon-type/application"
	"pokemon-lab-api/pkg/mderrors"

	"github.com/gofiber/fiber/v2"
)

type RetrieveAllPokemonTypeController struct {
	s application.RetrieveAllPokemonTypeUseCase
}

func (c *RetrieveAllPokemonTypeController) Handle(ctx *fiber.Ctx) error {

	pkmts, errR := c.s.Do()

	if errR != nil {
		return mderrors.NewMetadataError(errR)
	}

	return ctx.JSON(pkmts)
}

func NewRetrieveAllPokemonTypeController(s application.RetrieveAllPokemonTypeUseCase) *RetrieveAllPokemonTypeController {
	return &RetrieveAllPokemonTypeController{
		s: s,
	}
}
