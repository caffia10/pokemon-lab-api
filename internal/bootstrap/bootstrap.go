package bootstrap

import (
	"pokemon-lab-api/internal/http"

	"go.uber.org/dig"
)

func InitServer() {

	c := dig.New()

	ids := []any{}

	ids = append(ids, RetrieveInjections()...)

	for _, d := range ids {
		panic(c.Provide(d))
	}

	err := c.Invoke(http.StarListen)
	if err != nil {
		panic(err)
	}

}
