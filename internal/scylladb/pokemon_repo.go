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
)

type pokemonDto struct {
	Id       string   `db:"id"`
	Name     string   `db:"name"`
	Category string   `db:"category"`
	Weight   string   `db:"weight"`
	ImgUrl   string   `db:"img_url"`
	Types    []string `db:"types"` // save only the id
}

type pokemonRepository struct {
	session *gocqlx.Session
	// table allows for simple CRUD operations based on personMetadata.
	table *table.Table
}

func NewPokemonRepository(s *gocqlx.Session, c *config.Config) (core.PokemonRepository, error) {

	var pokemonMetadata = table.Metadata{
		Name:    c.PokemonScyllasTable,
		Columns: []string{"id", "name", "category", "weight", "img_url", "types"},
		PartKey: []string{"id"},
	}

	r := &pokemonRepository{
		session: s,
		table:   table.New(pokemonMetadata),
	}

	return r, r.applyTable()
}

func (r *pokemonRepository) RetrieveAll() ([]core.Pokemon, error) {

	var pkmDtos []*pokemonDto

	q := r.session.Query(fmt.Sprintf("SELECT * FROM %s", r.table.Name()), nil)

	if err := q.SelectRelease(&pkmDtos); err != nil {
		if errors.Is(err, gocql.ErrNotFound) {
			return nil, mderrors.NewMetadataError(core.ErrNotFoundPokemon)
		}
		return nil, mderrors.NewMetadataError(err)
	}

	pts := make([]core.Pokemon, len(pkmDtos))

	for i, dto := range pkmDtos {
		pts[i] = mapDtoToModel(dto)
	}

	return pts, nil
}

func (r *pokemonRepository) RetriveById(id string) (*core.Pokemon, error) {

	pkmDto := pokemonDto{Id: id}

	q := r.session.Query(r.table.Get()).BindStruct(pkmDto)

	if err := q.GetRelease(&pkmDto); err != nil {
		if errors.Is(err, gocql.ErrNotFound) {
			return nil, mderrors.NewMetadataError(
				core.ErrNotFoundPokemon,
				mderrors.MakeIdMetaData(id))
		}
		return nil, mderrors.NewMetadataError(
			err,
			mderrors.MakeIdMetaData(id))
	}

	pkm := mapDtoToModel(&pkmDto)

	return &pkm, nil
}

/*
func (r *PokemonScylla) RetriveByIds(ids []string) (*core.Pokemon, error) {

	lf := []zap.Field{
		zap.String("logger", "PokemonScylla"),
		zap.String("sub-logger", "RetriveByIds"),
	}

	stmt, names := qb.Select(r.table.Name()).Columns(r.table.Metadata().Columns...).Where(qb.In("id")).ToCql()

	r.logger.Info("retrieving pokemon", lf...)

	q := r.session.Query(r.table.Get()).BindStruct(pkmDto)

	if err := q.GetRelease(&pkmDto); err != nil {
		if errors.Is(err, gocql.ErrNotFound) {
			r.logger.Warn("no document was found")
			return nil, domain.ErrNotFoundPokemon
		}
		lf = append(lf, zap.NamedError("error", err))
		r.logger.Error("unexpected error at retriving pokemon", lf...)
		return nil, err
	}

	pkm := &domain.Pokemon{

		Id:       pkmDto.Id,
		Name:     pkmDto.Name,
		ImgUrl:   pkmDto.ImgUrl,
		Category: pkmDto.Category,
		Weight:   pkmDto.Weight,
		Types:    make([]*domain.PokemonType, len(pkmDto.Types)),
	}

	for _, v := range pkmDto.Types {
		pkm.Types = append(pkm.Types, &domain.PokemonType{Name: *v})
	}

	return pkm, nil
}*/

func (r *pokemonRepository) Create(pkm *core.Pokemon) error {
	q := r.session.Query(r.table.Insert()).BindStruct(initializeDtoFromModel(pkm))

	if err := q.ExecRelease(); err != nil {
		return mderrors.NewMetadataError(err, mderrors.Metadata{Key: "pokemon", Value: pkm})
	}

	return nil
}

func (r *pokemonRepository) CreateMany(pkms []core.Pokemon) error {

	errChan := make(chan error)
	var wg sync.WaitGroup
	for _, pkm := range pkms {
		wg.Add(1)
		go func(p *core.Pokemon, w *sync.WaitGroup) {
			defer w.Done()
			err := r.Create(p)
			if err != nil {
				errChan <- err
			}
		}(&pkm, &wg)
	}

	wg.Wait()
	close(errChan)

	if len(errChan) > 0 {

		me := NewBufferedMultiError(len(errChan))
		for err := range errChan {
			me.Append(err)
		}

		return mderrors.NewMetadataError(me)
	}

	return nil
}

func (r *pokemonRepository) Update(pkm *core.Pokemon) error {

	dto := initializeDtoFromModel(pkm)
	dto.Id = pkm.Id

	q := r.session.Query(r.table.Update(r.table.Metadata().Columns[1:]...)).BindStruct(dto)

	if err := q.ExecRelease(); err != nil {
		return mderrors.NewMetadataError(err, mderrors.Metadata{Key: "pokemon", Value: pkm})
	}

	return nil

}

// Delete is idempotent: in Scylla a DELETE writes a tombstone whether or not the row
// exists, so a missing id is not an error.
func (r *pokemonRepository) Delete(id string) error {

	q := r.session.Query(r.table.Delete()).BindStruct(pokemonDto{Id: id})

	if err := q.ExecRelease(); err != nil {
		return mderrors.NewMetadataError(err, mderrors.Metadata{Key: "pokemon_id", Value: id})
	}

	return nil
}

func initializeDtoFromModel(pkm *core.Pokemon) *pokemonDto {

	pkmd := &pokemonDto{
		Id:       uuid.NewString(),
		Name:     pkm.Name,
		Category: pkm.Category,
		Weight:   pkm.Weight,
		ImgUrl:   pkm.ImgUrl,
		Types:    make([]string, 0, len(pkm.Types)),
	}

	for _, t := range pkm.Types {
		pkmd.Types = append(pkmd.Types, t.Name)
	}

	return pkmd
}

func (r *pokemonRepository) applyTable() error {
	data := make([]any, len(r.table.Metadata().Columns)+1)
	data[0] = r.table.Name()
	for i := 0; i < len(r.table.Metadata().Columns); i++ {
		data[i+1] = r.table.Metadata().Columns[i]
	}
	query := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
		%s UUID PRIMARY KEY,
		%s TEXT,
		%s TEXT,
		%s TEXT,
		%s TEXT,
		%s LIST<TEXT>
	);`, data...)

	return r.session.ExecStmt(query)
}

func mapDtoToModel(dto *pokemonDto) core.Pokemon {
	tps := make([]core.PokemonType, len(dto.Types))

	for i, v := range dto.Types {
		tps[i] = core.PokemonType{Name: v}
	}

	return core.Pokemon{

		Id:       dto.Id,
		Name:     dto.Name,
		ImgUrl:   dto.ImgUrl,
		Category: dto.Category,
		Weight:   dto.Weight,
		Types:    tps,
	}

}
