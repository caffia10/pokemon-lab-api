package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	MongoUri                   string `envconfig:"MONGO_URI"`
	MongoDatabase              string `envconfig:"MONGO_DATABASE"`
	PokemonMongoCollection     string `envconfig:"POKEMON_MONGO_COLLECTION"`
	PokemonTypeMongoCollection string `envconfig:"POKEMON_TYPE_MONGO_COLLECTION"`
	ScyllaHost                 []string
	PokemonScyllasTable        string `envconfig:"POKEMON_SCYLLAS_TABLE"`
}

func NewConfig() *Config {
	cfg := &Config{}

	err := envconfig.Process("", cfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	return cfg
}
