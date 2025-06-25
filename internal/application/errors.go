package application

import (
	"errors"
	"pokemon-lab-api/internal/domain"
	"strings"
)

const (
	basicMsg = "Some pokemon type failt at their creation"
)

var (
	AllPokemonTypesFail = errors.New("all pokemon type failt at their creation")
)

type SomePokemonTypeFailCreationError struct {
	pkmst []*domain.PokemonType
}

func (e *SomePokemonTypeFailCreationError) Error() string {
	if len(e.pkmst) == 0 {
		return basicMsg
	}
	var sb strings.Builder
	sb.WriteString(basicMsg)
	sb.WriteString(", [")
	sb.WriteString(e.pkmst[0].Name)
	for i := 1; i < len(e.pkmst); i++ {
		sb.WriteRune(',')
		sb.WriteString(e.pkmst[i].Name)
	}
	sb.WriteRune(']')
	return sb.String()
}

func NewSomePokemonFailCreationError(pkmst []*domain.PokemonType) error {

	return &SomePokemonTypeFailCreationError{
		pkmst: pkmst,
	}
}
