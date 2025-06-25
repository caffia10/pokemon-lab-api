package domain

import "github.com/pkg/errors"

var ErrNotFoundPokemon = errors.New("not found pokemon")
var ErrNotFoundPokemonType = errors.New("not found pokemon type")
