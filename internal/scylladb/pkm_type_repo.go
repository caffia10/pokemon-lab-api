package scylladb

import (
	"errors"
	"fmt"
	"pokemon-lab-api/internal/config"
	"pokemon-lab-api/internal/core"
	"pokemon-lab-api/pkg/mderrors"
	"sync"

	"github.com/gocql/gocql"
	"github.com/google/uuid"
	"github.com/scylladb/gocqlx/table"
	"github.com/scylladb/gocqlx/v3"
	"go.uber.org/zap"
)

var ErrAllPokemonTypesFail = errors.New("all pokemon types fail to be created")
var ErrSomePokemonTypesFail = errors.New("some pokemon types fail to be created")

type pokemonTypeDto struct {
	Id   string `db:"id"`
	Name string `db:"name"`
}

type pokemonTypeRepository struct {
	session *gocqlx.Session
	// table allows for simple CRUD operations based on personMetadata.
	table *table.Table

	logger *zap.Logger
}

func NewPokemonTypeRepository(s *gocqlx.Session, c *config.Config, l *zap.Logger) (core.PokemonTypeRepository, error) {

	var pokemonMetadata = table.Metadata{
		Name:    c.PokemonTypeScyllasTable,
		Columns: []string{"id", "name"},
		PartKey: []string{"id"},
		SortKey: []string{"id"},
	}

	r := &pokemonTypeRepository{
		session: s,
		table:   table.New(pokemonMetadata),
		logger:  l,
	}

	return r, r.applyTable()
}

func (r *pokemonTypeRepository) RetrieveAll() ([]core.PokemonType, error) {

	var pkmDtos []pokemonTypeDto

	q := r.session.Query(fmt.Sprintf("SELECT * FROM %s", r.table.Name()), nil)

	if err := q.SelectRelease(&pkmDtos); err != nil {
		if errors.Is(err, gocql.ErrNotFound) {
			return nil, core.ErrNotFoundPokemonType
		}
		return nil, err
	}

	pts := make([]core.PokemonType, len(pkmDtos))

	for i, dto := range pkmDtos {
		pts[i] = core.PokemonType(dto)
	}

	return pts, nil
}

func (r *pokemonTypeRepository) Create(pkmt *core.PokemonType) error {

	r.logger.Debug("creating pokemon type on scyllas", zap.String("pokemon-type-name", pkmt.Name))

	q := r.session.Query(r.table.Insert()).BindStruct(&pokemonTypeDto{
		Id:   uuid.NewString(),
		Name: pkmt.Name,
	})

	if err := q.ExecRelease(); err != nil {
		return err
	}

	return nil
}

func (r *pokemonTypeRepository) CreateMany(pkmts []core.PokemonType) error {

	var wg sync.WaitGroup
	errChannel := make(chan core.PokemonType)
	for _, pkmt := range pkmts {
		wg.Add(1)
		go func(pt *core.PokemonType) {
			defer wg.Done()
			err := r.Create(pt)
			if err != nil {
				errChannel <- *pt
			}
		}(&pkmt)
	}
	wg.Wait()
	close(errChannel)

	if len(errChannel) > 0 {

		if len(errChannel) == len(pkmts) {
			return ErrAllPokemonTypesFail
		}

		ept := make([]mderrors.Metadata, len(errChannel))
		for pt := range errChannel {

			ept = append(ept, mderrors.Metadata{Key: "failed-pokemon-type", Value: pt})

		}

		return mderrors.NewMetadataError(ErrSomePokemonTypesFail, ept...)
	}

	return nil
}

func (r *pokemonTypeRepository) applyTable() error {
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
