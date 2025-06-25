package fiberapi

import (
	"pokemon-lab-api/internal/infrastructure/controller"

	"github.com/gofiber/fiber/v2"
)

type restServer struct {
	app         *fiber.App
	cpCtrl      *controller.CreatePokemon
	cptCtrl     *controller.CreatePokemonType
	rtapkmtCtrl *controller.RetrieveAllPokemonType
}

func StarListen(rs *restServer) {
	rs.app.Listen(":3030")
}

func NewFiberApiServer(
	cpCtrl *controller.CreatePokemon,
	cptCtrl *controller.CreatePokemonType,
	rtapkmtCtrl *controller.RetrieveAllPokemonType) *restServer {
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
