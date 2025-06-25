package server

import (
	"go.uber.org/dig"
)

func InitServer() {

	c := dig.New()

	for _, d := range RetrieveInfraInjections() {
		c.Provide(d)
	}

	err := c.Invoke(StarListen)
	if err != nil {
		panic(err)
	}

}
