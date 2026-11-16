package http

import (
	"errors"
	"pokemon-lab-api/internal/core"
	"pokemon-lab-api/pkg/mderrors"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type pokemonRequest struct {
	Id       string               `json:"id"`
	Name     string               `json:"name"`
	Weight   string               `json:"weight"`
	Category string               `json:"category"`
	ImgUrl   string               `json:"imgUrl"`
	Types    []pokemonTypeRequest `json:"types"`
}

func (req *pokemonRequest) toPokemon() *core.Pokemon {
	pkm := &core.Pokemon{
		Id:       req.Id,
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

// parseBody parses the request body and keeps the parser's error detail
// instead of masking it behind a generic 400.
func parseBody(ctx *fiber.Ctx, out any) error {
	err := ctx.BodyParser(out)
	if err == nil {
		return nil
	}

	var fe *fiber.Error
	if errors.As(err, &fe) {
		// e.g. 422 when the Content-Type is missing or not supported
		return fiber.NewError(fe.Code, fe.Message+": expected Content-Type application/json, got '"+string(ctx.Request().Header.ContentType())+"'")
	}

	return fiber.NewError(fiber.StatusBadRequest, "invalid request body: "+err.Error())
}

func makeCreatePokemonHandler(repo core.PokemonRepository) fiber.Handler {

	return func(ctx *fiber.Ctx) error {

		req := new(pokemonRequest)
		if err := parseBody(ctx, req); err != nil {
			return err
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
		if err := parseBody(ctx, req); err != nil {
			return err
		}

		if req.Id == "" {
			return fiber.NewError(fiber.StatusBadRequest, "id is required to update a pokemon")
		}

		pkm := req.toPokemon()

		errC := repo.Update(pkm)

		if errC != nil {
			return mderrors.NewMetadataError(errC)
		}

		return ctx.SendStatus(fiber.StatusOK)
	}
}

func makeDeletePokemonHandler(repo core.PokemonRepository) fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		id := ctx.Params("id")

		// the partition key is a UUID: reject anything else here instead of letting the
		// driver fail to marshal it and surface as a 409
		if _, err := uuid.Parse(id); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "id must be a UUID: "+id)
		}

		if err := repo.Delete(id); err != nil {
			return mderrors.NewMetadataError(err)
		}

		return ctx.SendStatus(fiber.StatusOK)
	}
}
