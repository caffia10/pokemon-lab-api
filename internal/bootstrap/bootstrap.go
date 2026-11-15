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
		if err := c.Provide(d); err != nil {
			panic(err)
		}
	}

	err := c.Invoke(http.StarListen)
	if err != nil {
		panic(err)
	}

}
