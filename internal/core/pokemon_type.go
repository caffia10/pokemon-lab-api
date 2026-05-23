package core

type PokemonType struct {
	Id   string
	Name string
}

type PokemonTypeRepository interface {
	RetrieveAll() ([]PokemonType, error)
	Create(pkmt *PokemonType) error
	CreateMany(pkmts []PokemonType) error
}
