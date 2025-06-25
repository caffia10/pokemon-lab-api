package domain

type Evolution struct {
	Id          string
	Requirement string
	Pokemon     *Pokemon
}

type PokemonType struct {
	Id   string
	Name string
}

type Pokemon struct {
	Id       string
	Name     string
	Weight   string
	Category string
	ImgUrl   string
	Types    []PokemonType
	// Evolutions hold the pokemon evolutions. This data is lazy loading
	Evolutions []Evolution
}

type PokemonRepository interface {
	RetriveById(id string) (*Pokemon, error)
	Create(pkm *Pokemon) error
	CreateManyPokemon(pkms []Pokemon) error
}

type PokemonTypeRepository interface {
	RetriveAll() ([]PokemonType, error)
	Create(pkmt *PokemonType) error
}
