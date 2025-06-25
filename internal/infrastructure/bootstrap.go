package infrastructure

import (
	"pokemon-lab-api/internal/infrastructure/fiberapi"
	"pokemon-lab-api/internal/infrastructure/injections"

	"go.uber.org/dig"
)

func InitServer() {

	c := dig.New()

	ids := []interface{}{}

	ids = append(ids, injections.RetrieveInfraInjections()...)
	ids = append(ids, injections.RetrieveDomainInjections()...)
	ids = append(ids, injections.RetrieveAppInjections()...)

	for _, d := range ids {
		c.Provide(d)
	}

	err := c.Invoke(fiberapi.StarListen)
	if err != nil {
		panic(err)
	}

}
