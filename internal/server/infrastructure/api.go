package infrastructure

import (
	pkmInfra "pokemon-lab-api/internal/pokemon/infrastructure/http"

	"github.com/gofiber/fiber/v2"
)

type restServer struct {
	app    *fiber.App
	cpCtrl *pkmInfra.CreatePokemonController
}

func StarListen(rs *restServer) {
	rs.app.Listen(":3030")
}

func NewFiberApiServer(cpCtrl *pkmInfra.CreatePokemonController) *restServer {
	app := fiber.New()

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("Pong")
	})

	app.Post("/create-pokemon", cpCtrl.Handle)

	return &restServer{
		app:    app,
		cpCtrl: cpCtrl,
	}
}
