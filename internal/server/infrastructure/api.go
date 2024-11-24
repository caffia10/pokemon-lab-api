package infrastructure

import (
	pkmtInfra "pokemon-lab-api/internal/pokemon-type/infrastructure/http"
	pkmInfra "pokemon-lab-api/internal/pokemon/infrastructure/http"

	"github.com/gofiber/fiber/v2"
)

type restServer struct {
	app         *fiber.App
	cpCtrl      *pkmInfra.CreatePokemonController
	cptCtrl     *pkmtInfra.CreatePokemonTypeController
	rtapkmtCtrl *pkmtInfra.RetrieveAllPokemonTypeController
}

func StarListen(rs *restServer) {
	rs.app.Listen(":3030")
}

func NewFiberApiServer(
	cpCtrl *pkmInfra.CreatePokemonController,
	cptCtrl *pkmtInfra.CreatePokemonTypeController,
	rtapkmtCtrl *pkmtInfra.RetrieveAllPokemonTypeController) *restServer {
	app := fiber.New()

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("Pong")
	})

	api := app.Group("api")
	api.Post("/pokemons", cpCtrl.Validate, cpCtrl.Handle)
	api.Post("/pokemon-types", cptCtrl.Validate, cptCtrl.Handle)
	api.Get("/pokemon-types", rtapkmtCtrl.Handle)

	return &restServer{
		app:         app,
		cpCtrl:      cpCtrl,
		cptCtrl:     cptCtrl,
		rtapkmtCtrl: rtapkmtCtrl,
	}
}
