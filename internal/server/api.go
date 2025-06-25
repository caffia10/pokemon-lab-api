package server

import (
	"pokemon-lab-api/internal/pokemon"
	pokemontype "pokemon-lab-api/internal/pokemon-type"

	"github.com/gofiber/fiber/v2"
)

type restServer struct {
	app         *fiber.App
	cpCtrl      *pokemon.CreatePokemonController
	cptCtrl     *pokemontype.CreatePokemonTypeController
	rtapkmtCtrl *pokemontype.RetrieveAllPokemonTypeController
}

func StarListen(rs *restServer) {
	rs.app.Listen(":3030")
}

func NewFiberApiServer(
	cpCtrl *pokemon.CreatePokemonController,
	cptCtrl *pokemontype.CreatePokemonTypeController,
	rtapkmtCtrl *pokemontype.RetrieveAllPokemonTypeController) *restServer {
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
