package http

import (
	"pokemon-lab-api/internal/core"
	"pokemon-lab-api/pkg/mderrors"

	"github.com/gofiber/fiber/v2"
)

type pokemonRequest struct {
	Name     string               `json:"name"`
	Weight   string               `json:"weight"`
	Category string               `json:"category"`
	ImgUrl   string               `json:"img_url"`
	Types    []pokemonTypeRequest `json:"types"`
}

func (req *pokemonRequest) toPokemon() *core.Pokemon {
	pkm := &core.Pokemon{
		Name:     req.Name,
		Category: req.Category,
		Weight:   req.Weight,
		ImgUrl:   req.ImgUrl,
		Types:    make([]core.PokemonType, 0, len(req.Types)),
	}

	for _, v := range req.Types {
		pkm.Types = append(pkm.Types, (core.PokemonType)(v))
	}

	return pkm
}

func makeCreatePokemonHandler(repo core.PokemonRepository) fiber.Handler {

	return func(ctx *fiber.Ctx) error {

		req := new(pokemonRequest)
		errB := ctx.BodyParser(req)

		if errB != nil {
			return fiber.ErrBadRequest
		}

		pkm := req.toPokemon()

		errC := repo.Create(pkm)

		if errC != nil {
			return mderrors.NewMetadataError(errC)
		}

		return ctx.SendStatus(fiber.StatusOK)
	}
}

func makeRetrieveAllPokemonHandler(repo core.PokemonRepository) fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		pkmts, errR := repo.RetrieveAll()

		if errR != nil {
			return mderrors.NewMetadataError(errR)
		}

		return ctx.JSON(pkmts)
	}
}

func makeRetrievePokemonHandler(repo core.PokemonRepository) fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		id := ctx.Params("id")

		pkmts, errR := repo.RetriveById(id)

		if errR != nil {
			return mderrors.NewMetadataError(errR)
		}

		return ctx.JSON(pkmts)
	}
}

func makeUpdatePokemonHandler(repo core.PokemonRepository) fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		req := new(pokemonRequest)
		errB := ctx.BodyParser(req)

		if errB != nil {
			return fiber.ErrBadRequest
		}

		pkm := req.toPokemon()

		errC := repo.Update(pkm)

		if errC != nil {
			return mderrors.NewMetadataError(errC)
		}

		return ctx.SendStatus(fiber.StatusOK)
	}
}
