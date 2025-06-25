package pokemon

import pokemontype "pokemon-lab-api/internal/pokemon-type"

type PokemonType = pokemontype.PokemonType

type Pokemon struct {
	Id       string
	Name     string
	Weight   string
	Category string
	ImgUrl   string
	Types    []*PokemonType
}

type PokemonRepository interface {
	RetriveById(id string) (*Pokemon, error)
	Create(pkm *Pokemon) error
	CreateManyPokemon(pkms []*Pokemon) error
}
