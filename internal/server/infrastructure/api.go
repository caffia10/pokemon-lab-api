package infrastructure

import (
	"fmt"
	pkmtInfra "pokemon-lab-api/internal/pokemon-type/infrastructure/http"
	pkmInfra "pokemon-lab-api/internal/pokemon/infrastructure/http"
	"pokemon-lab-api/internal/server/infrastructure/config"
	"pokemon-lab-api/internal/server/infrastructure/middlewares"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type restServer struct {
	listenPort  string
	app         *fiber.App
	cpCtrl      *pkmInfra.CreatePokemonController
	cptCtrl     *pkmtInfra.CreatePokemonTypeController
	rtapkmtCtrl *pkmtInfra.RetrieveAllPokemonTypeController
	rtapkmCtrl  *pkmInfra.RetrieveAllPokemonController
}

func StarListen(rs *restServer) {
	rs.app.Listen(rs.listenPort)
}

func NewFiberApiServer(
	cfg *config.Config,
	logger *zap.Logger,
	cpCtrl *pkmInfra.CreatePokemonController,
	cptCtrl *pkmtInfra.CreatePokemonTypeController,
	rtapkmtCtrl *pkmtInfra.RetrieveAllPokemonTypeController,
	rtapkmCtrl *pkmInfra.RetrieveAllPokemonController) *restServer {
	app := fiber.New(fiber.Config{
		ErrorHandler: middlewares.MakeErrorHandler(logger),
	})

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("Pong")
	})

	api := app.Group("api")
	api.Post("/pokemons", cpCtrl.Validate, cpCtrl.Handle)
	api.Get("/pokemons", rtapkmCtrl.Handle)
	api.Post("/pokemon-types", cptCtrl.Validate, cptCtrl.Handle)
	api.Get("/pokemon-types", rtapkmtCtrl.Handle)

	return &restServer{
		listenPort:  fmt.Sprintf(":%s", cfg.ServerPort),
		app:         app,
		cpCtrl:      cpCtrl,
		cptCtrl:     cptCtrl,
		rtapkmtCtrl: rtapkmtCtrl,
		rtapkmCtrl:  rtapkmCtrl,
	}
}
