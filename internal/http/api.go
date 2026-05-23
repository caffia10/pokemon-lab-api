package http

import (
	"fmt"
	"pokemon-lab-api/internal/config"
	"pokemon-lab-api/internal/core"
	"pokemon-lab-api/internal/http/middlewares"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type restServer struct {
	listenPort string
	app        *fiber.App
}

func StarListen(rs *restServer) error {
	return rs.app.Listen(rs.listenPort)
}

func NewFiberApiServer(
	cfg *config.Config,
	logger *zap.Logger,
	pokemonRepository core.PokemonRepository,
	pokemonTypeRepository core.PokemonTypeRepository) *restServer {
	app := fiber.New(fiber.Config{
		ErrorHandler: middlewares.MakeErrorHandler(logger),
	})

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("Pong")
	})

	api := app.Group("api")
	api.Post("/pokemons", makeCreatePokemonHandler(pokemonRepository))
	api.Put("/pokemons", makeUpdatePokemonHandler(pokemonRepository))
	api.Get("/pokemons", makeRetrieveAllPokemonHandler(pokemonRepository))
	api.Get("/pokemons/:id", makeRetrievePokemonHandler(pokemonRepository))
	api.Post("/pokemon-types", makeCreatePokemonTypeHandler(pokemonTypeRepository))
	api.Get("/pokemon-types", makeRetrieveAllPokemonTypeHandler(pokemonTypeRepository))

	return &restServer{
		listenPort: fmt.Sprintf(":%s", cfg.ServerPort),
		app:        app,
	}
}
