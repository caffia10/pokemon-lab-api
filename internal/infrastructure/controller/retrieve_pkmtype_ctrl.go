package controller

import (
	"pokemon-lab-api/internal/application"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type RetrieveAllPokemonType struct {
	s      application.RetrieveAllPokemonTypeUseCase
	logger *zap.Logger
}

func (c *RetrieveAllPokemonType) Handle(ctx *fiber.Ctx) error {

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

func NewRetrieveAllPokemonTypeController(s application.RetrieveAllPokemonTypeUseCase, logger *zap.Logger) *RetrieveAllPokemonType {
	return &RetrieveAllPokemonType{
		s:      s,
		logger: logger,
	}
}
