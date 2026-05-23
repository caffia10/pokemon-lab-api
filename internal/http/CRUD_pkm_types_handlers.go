package http

import (
	"pokemon-lab-api/internal/core"
	"pokemon-lab-api/pkg/mderrors"

	"github.com/gofiber/fiber/v2"
)

func makeCreatePokemonTypeHandler(s core.PokemonTypeRepository) fiber.Handler {

	return func(ctx *fiber.Ctx) error {

		req := new(pokemonTypeRequest)
		errB := ctx.BodyParser(req)

		if errB != nil {
			return fiber.ErrBadRequest
		}

		errC := s.Create((*core.PokemonType)(req))

		if errC != nil {
			return fiber.ErrConflict
		}

		return ctx.SendStatus(fiber.StatusOK)
	}
}

func makeRetrieveAllPokemonTypeHandler(s core.PokemonTypeRepository) fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		pkmts, errR := s.RetrieveAll()

		if errR != nil {
			return mderrors.NewMetadataError(errR)
		}

		return ctx.JSON(pkmts)
	}
}
