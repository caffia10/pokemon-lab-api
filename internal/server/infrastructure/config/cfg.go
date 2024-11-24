package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	MongoUri                   string   `envconfig:"MONGO_URI"`
	MongoDatabase              string   `envconfig:"MONGO_DATABASE"`
	PokemonMongoCollection     string   `envconfig:"POKEMON_MONGO_COLLECTION"`
	PokemonTypeMongoCollection string   `envconfig:"POKEMON_TYPE_MONGO_COLLECTION"`
	ScyllaHosts                []string `envconfig:"SCYLLA_HOSTS"`
	PokemonScyllasTable        string   `envconfig:"POKEMON_SCYLLAS_TABLE"`
	PokemonTypeScyllasTable    string   `envconfig:"POKEMON_TYPE_SCYLLAS_TABLE"`
	ScyllaKeySpaces            string   `envconfig:"SCYLLA_KEY_SPACE"`
}

func NewConfig() *Config {
	cfg := &Config{}

	err := envconfig.Process("", cfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	return cfg
}
