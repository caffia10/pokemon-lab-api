package pokemontype

import (
	"errors"
	"fmt"
	"pokemon-lab-api/internal/config"

	"github.com/gocql/gocql"
	"github.com/google/uuid"
	"github.com/scylladb/gocqlx/table"
	"github.com/scylladb/gocqlx/v3"
	"go.uber.org/zap"
)

type PokemonTypeScyllasRepository struct {
	session *gocqlx.Session
	// table allows for simple CRUD operations based on personMetadata.
	table *table.Table

	logger *zap.Logger
}

func (r *PokemonTypeScyllasRepository) RetriveAll() ([]*PokemonType, error) {

	lf := []zap.Field{
		zap.String("logger", "PokemonType"),
		zap.String("sub-logger", "RetriveAll"),
	}

	var pkmDtos []*pokemonTypeDto

	r.logger.Info("retrieving pokemon types", lf...)

	q := r.session.Query(fmt.Sprintf("SELECT * FROM %s", r.table.Name()), nil)

	if err := q.SelectRelease(&pkmDtos); err != nil {
		if errors.Is(err, gocql.ErrNotFound) {
			r.logger.Warn("no document was found")
			return nil, ErrNotFoundPokemonType
		}
		lf = append(lf, zap.NamedError("error", err))
		r.logger.Error("unexpected error at retriving pokemon types", lf...)
		return nil, err
	}

	pts := make([]*PokemonType, len(pkmDtos))

	for i, dto := range pkmDtos {
		pts[i] = (*PokemonType)(dto)
	}

	return pts, nil
}

func (r *PokemonTypeScyllasRepository) Create(pkmt *PokemonType) error {

	lf := []zap.Field{
		zap.String("logger", "PokemonType"),
		zap.String("sub-logger", "Create"),
		zap.String("pokemon-type-name", pkmt.Name),
	}

	r.logger.Info("creating pokemon type", lf...)

	q := r.session.Query(r.table.Insert()).BindStruct(&pokemonTypeDto{
		Id:   uuid.NewString(),
		Name: pkmt.Name,
	})

	if err := q.ExecRelease(); err != nil {
		lf = append(lf, zap.NamedError("error", err))
		r.logger.Error("unexpected error at creating pokemon type", lf...)

		return err
	}

	return nil
}

func (r *PokemonTypeScyllasRepository) applyTable() error {
	data := make([]any, len(r.table.Metadata().Columns)+1)
	data[0] = r.table.Name()
	for i := 0; i < len(r.table.Metadata().Columns); i++ {
		data[i+1] = r.table.Metadata().Columns[i]
	}
	query := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
		%s UUID PRIMARY KEY,
		%s TEXT,
	);`, data...)

	return r.session.ExecStmt(query)
}

func NewPokemonTypeRepository(s *gocqlx.Session, c *config.Config, l *zap.Logger) (PokemonTypeRepository, error) {

	var pokemonMetadata = table.Metadata{
		Name:    c.PokemonTypeScyllasTable,
		Columns: []string{"id", "name"},
		PartKey: []string{"id"},
		SortKey: []string{"id"},
	}

	r := &PokemonTypeScyllasRepository{
		session: s,
		table:   table.New(pokemonMetadata),
		logger:  l,
	}

	return r, r.applyTable()
}
