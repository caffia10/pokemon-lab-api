package http

import (
	"pokemon-lab-api/internal/pokemon/application"
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

func (req *pokemonRequest) toPokemon() *application.Pokemon {
	pkm := &application.Pokemon{
		Name:     req.Name,
		Category: req.Category,
		Weight:   req.Weight,
		ImgUrl:   req.ImgUrl,
		Types:    make([]application.PokemonType, 0, len(req.Types)),
	}

	for _, v := range req.Types {
		pkm.Types = append(pkm.Types, (application.PokemonType)(v))
	}

	return pkm
}

func MakeCreatePokemonHandler(s application.CreatePokemonUseCase) fiber.Handler {

	return func(ctx *fiber.Ctx) error {

		req := new(pokemonRequest)
		errB := ctx.BodyParser(req)

		if errB != nil {
			return fiber.ErrBadRequest
		}

		pkm := req.toPokemon()

		errC := s.Do(pkm)

		if errC != nil {
			return mderrors.NewMetadataError(errC)
		}

		return ctx.SendStatus(fiber.StatusOK)
	}
}

func MakeRetrieveAllPokemonHandler(s application.RetrieveAllPokemonUseCase) fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		pkmts, errR := s.Do()

		if errR != nil {
			return mderrors.NewMetadataError(errR)
		}

		return ctx.JSON(pkmts)
	}
}

func MakeRetrievePokemonHandler(s application.RetrievePokemonbyIdUseCase) fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		id := ctx.Params("id")

		pkmts, errR := s.Do(id)

		if errR != nil {
			return mderrors.NewMetadataError(errR)
		}

		return ctx.JSON(pkmts)
	}
}

func MakeUpdatePokemonHandler(s application.UpdatePokemonUseCase) fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		req := new(pokemonRequest)
		errB := ctx.BodyParser(req)

		if errB != nil {
			return fiber.ErrBadRequest
		}

		pkm := req.toPokemon()

		errC := s.Do(pkm)

		if errC != nil {
			return mderrors.NewMetadataError(errC)
		}

		return ctx.SendStatus(fiber.StatusOK)
	}
}
