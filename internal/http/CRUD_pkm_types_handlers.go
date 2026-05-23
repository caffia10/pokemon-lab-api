package http

import (
	"pokemon-lab-api/internal/pokemon-type/application"
	"pokemon-lab-api/pkg/mderrors"

	"github.com/gofiber/fiber/v2"
)

func MakeCreatePokemonTypeHandler(s application.CreatePokemonTypeUseCase) fiber.Handler {

	return func(ctx *fiber.Ctx) error {

		req := new(pokemonTypeRequest)
		errB := ctx.BodyParser(req)

		if errB != nil {
			return fiber.ErrBadRequest
		}

		errC := s.Do((*application.PokemonType)(req))

		if errC != nil {
			return fiber.ErrConflict
		}

		return ctx.SendStatus(fiber.StatusOK)
	}
}

func MakeRetrieveAllPokemonTypeHandler(s application.RetrieveAllPokemonTypeUseCase) fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		pkmts, errR := s.Do()

		if errR != nil {
			return mderrors.NewMetadataError(errR)
		}

		return ctx.JSON(pkmts)
	}
}
