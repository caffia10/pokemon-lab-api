package server

import (
	serverInfra "pokemon-lab-api/internal/server/infrastructure"
	"pokemon-lab-api/internal/server/infrastructure/injections"

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

	err := c.Invoke(serverInfra.StarListen)
	if err != nil {
		panic(err)
	}

}
