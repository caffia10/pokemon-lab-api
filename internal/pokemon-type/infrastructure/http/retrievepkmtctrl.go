package http

import (
	"pokemon-lab-api/internal/pokemon-type/application"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type RetrieveAllPokemonTypeController struct {
	s      application.RetrieveAllPokemonTypeUseCase
	logger *zap.Logger
}

func (c *RetrieveAllPokemonTypeController) Handle(ctx *fiber.Ctx) error {

	lf := []zap.Field{
		zap.String("logger", "RetrieveAllPokemonTypeController"),
		zap.String("sub-logger", "Handle"),
	}

	pkmts, errR := c.s.Do()

	if errR != nil {
		lf = append(lf, zap.NamedError("error", errR))
		c.logger.Error("error at retrieving all pokemon type from service", lf...)
		return fiber.ErrConflict
	}

	return ctx.JSON(pkmts)
}

func NewRetrieveAllPokemonTypeController(s application.RetrieveAllPokemonTypeUseCase, logger *zap.Logger) *RetrieveAllPokemonTypeController {
	return &RetrieveAllPokemonTypeController{
		s:      s,
		logger: logger,
	}
}
