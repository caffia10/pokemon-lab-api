package pokemontype

type PokemonType struct {
	Id   string
	Name string
}

type PokemonTypeRepository interface {
	RetriveAll() ([]*PokemonType, error)
	Create(pkmt *PokemonType) error
}
