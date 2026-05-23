package domain

import "github.com/pkg/errors"

var (
	ErrNotFoundPokemon     = errors.New("not found pokemon")
	ErrNotFoundPokemonType = errors.New("not found pokemon type")
)
