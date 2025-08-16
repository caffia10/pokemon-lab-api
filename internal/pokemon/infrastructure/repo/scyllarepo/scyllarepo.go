package scyllarepo

import (
	"errors"
	"fmt"
	"pokemon-lab-api/internal/pokemon/domain"
	"pokemon-lab-api/internal/server/infrastructure/config"
	"pokemon-lab-api/pkg/mderrors"

	"sync"

	"github.com/gocql/gocql"
	"github.com/google/uuid"
	"github.com/scylladb/gocqlx/table"
	"github.com/scylladb/gocqlx/v3"
)

type PokemonScylla struct {
	session *gocqlx.Session
	// table allows for simple CRUD operations based on personMetadata.
	table *table.Table
}

func mapDtoToModel(dto *pokemonDto) domain.Pokemon {
	tps := make([]domain.PokemonType, len(dto.Types))

	for i, v := range dto.Types {
		tps[i] = domain.PokemonType{Name: v}
	}

	return domain.Pokemon{

		Id:       dto.Id,
		Name:     dto.Name,
		ImgUrl:   dto.ImgUrl,
		Category: dto.Category,
		Weight:   dto.Weight,
		Types:    tps,
	}

}

func (r *PokemonScylla) RetrieveAll() ([]domain.Pokemon, error) {

	var pkmDtos []*pokemonDto

	q := r.session.Query(fmt.Sprintf("SELECT * FROM %s", r.table.Name()), nil)

	if err := q.SelectRelease(&pkmDtos); err != nil {
		if errors.Is(err, gocql.ErrNotFound) {
			return nil, mderrors.NewMetadataError(domain.ErrNotFoundPokemon)
		}
		return nil, mderrors.NewMetadataError(err)
	}

	pts := make([]domain.Pokemon, len(pkmDtos))

	for i, dto := range pkmDtos {
		pts[i] = mapDtoToModel(dto)
	}

	return pts, nil
}

func (r *PokemonScylla) RetriveById(id string) (*domain.Pokemon, error) {

	pkmDto := pokemonDto{Id: id}

	q := r.session.Query(r.table.Get()).BindStruct(pkmDto)

	if err := q.GetRelease(&pkmDto); err != nil {
		if errors.Is(err, gocql.ErrNotFound) {
			return nil, mderrors.NewMetadataError(
				domain.ErrNotFoundPokemon,
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
func (r *PokemonScylla) RetriveByIds(ids []string) (*domain.Pokemon, error) {

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

func initializeDtoFromModel(pkm *domain.Pokemon) *pokemonDto {

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

func (r *PokemonScylla) Create(pkm *domain.Pokemon) error {
	q := r.session.Query(r.table.Insert()).BindStruct(initializeDtoFromModel(pkm))

	if err := q.ExecRelease(); err != nil {
		return mderrors.NewMetadataError(err, mderrors.Metadata{Key: "pokemon", Value: pkm})
	}

	return nil
}

func (r *PokemonScylla) CreateManyPokemon(pkms []domain.Pokemon) error {

	errChan := make(chan error)
	var wg sync.WaitGroup
	for _, pkm := range pkms {
		wg.Add(1)
		go func(p *domain.Pokemon, w *sync.WaitGroup) {
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

func (r *PokemonScylla) applyTable() error {
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

func NewPokemonRepository(s *gocqlx.Session, c *config.Config) (domain.PokemonRepository, error) {

	var pokemonMetadata = table.Metadata{
		Name:    c.PokemonScyllasTable,
		Columns: []string{"id", "name", "category", "weight", "img_url", "types"},
		PartKey: []string{"first_name"},
		SortKey: []string{"last_name"},
	}

	r := &PokemonScylla{
		session: s,
		table:   table.New(pokemonMetadata),
	}

	return r, r.applyTable()
}
