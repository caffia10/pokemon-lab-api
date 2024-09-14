package domain

type PokemonType struct {
	Id   string
	Name string
}

type PokemonTypeRepository interface {
	RetriveById(id string) (*PokemonType, error)
	RetriveByIds(id []string) ([]*PokemonType, error)
	Create(pkmt *PokemonType) error
}
