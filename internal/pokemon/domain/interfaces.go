package domain

type RetrievePokemonbyIdUsecase interface {
	Do(id string) (*Pokemon, error)
}

type CreatePokemonUsecase interface {
	Do(pkm *Pokemon) error
}

type CreatePokemonInBulkUsecase interface {
	Do(pkms []*Pokemon) error
}
