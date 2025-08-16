package domain

import (
	pkmtDomain "pokemon-lab-api/internal/pokemon-type/domain"
)

type PokemonType = pkmtDomain.PokemonType

type Pokemon struct {
	Id       string
	Name     string
	Weight   string
	Category string
	ImgUrl   string
	Types    []PokemonType
}

type PokemonRepository interface {
	RetriveById(id string) (*Pokemon, error)
	RetrieveAll() ([]Pokemon, error)
	Create(pkm *Pokemon) error
	CreateManyPokemon(pkms []Pokemon) error
}
