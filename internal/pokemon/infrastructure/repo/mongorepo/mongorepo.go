package mongorepo

import (
	"context"
	"pokemon-lab-api/internal/pokemon/domain"
	"pokemon-lab-api/internal/server/infrastructure/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type PokemonMongo struct {
	coll   *mongo.Collection
	logger *zap.Logger
}

func (r *PokemonMongo) RetriveById(id string) (*domain.Pokemon, error) {

	lf := []zap.Field{
		zap.String("logger", "PokemonMongo"),
		zap.String("sub-logger", "RetriveById"),
		zap.String("pokemon-id", id),
	}

	r.logger.Info("retrieving pokemon", lf...)

	oId, errOId := primitive.ObjectIDFromHex(id)

	if errOId != nil {
		lf = append(lf, zap.NamedError("error", errOId))
		r.logger.Error("unexpected error at creating the objectId", lf...)
		return nil, errOId
	}

	pkmDto := new(pokemonDto)
	err := r.coll.FindOne(context.TODO(), bson.D{{Key: "_id", Value: oId}}).
		Decode(pkmDto)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			r.logger.Warn("no document was found")
			return nil, domain.ErrNotFoundPokemon
		}
		lf = append(lf, zap.NamedError("error", err))
		r.logger.Error("unexpected error at retriving pokemon", lf...)
		return nil, err
	}

	pkm := &domain.Pokemon{

		Id:       pkmDto.Id.String(),
		Name:     pkmDto.Name,
		ImgUrl:   pkmDto.ImgUrl,
		Category: pkmDto.Category,
		Weight:   pkmDto.Weight,
		Types:    make([]*domain.PokemonType, len(pkmDto.Types)),
	}

	for _, v := range pkmDto.Types {
		pkm.Types = append(pkm.Types, &domain.PokemonType{Id: v.String()})
	}

	return pkm, nil
}

func (r *PokemonMongo) initializeDtoFromModel(pkm *domain.Pokemon) *pokemonDto {

	pkmd := &pokemonDto{
		Id:       primitive.NewObjectID(),
		Name:     pkm.Name,
		Category: pkm.Category,
		Weight:   pkm.Weight,
		ImgUrl:   pkm.ImgUrl,
		Types:    make([]*primitive.ObjectID, len(pkm.Types)),
	}

	for _, v := range pkm.Types {
		oId, errOId := primitive.ObjectIDFromHex(v.Id)
		if errOId != nil {
			lf := []zap.Field{
				zap.String("logger", "PokemonMongo"),
				zap.String("sub-logger", "initializeDtoFromModel"),
				zap.String("pokemon-name", pkmd.Name),
				zap.NamedError("error", errOId),
			}

			r.logger.Warn("the pokemon can't be created", lf...)
			continue
		}
		pkmd.Types = append(pkmd.Types, &oId)
	}

	return pkmd
}

func (r *PokemonMongo) Create(pkm *domain.Pokemon) error {

	lf := []zap.Field{
		zap.String("logger", "PokemonMongo"),
		zap.String("sub-logger", "CreatePokemon"),
		zap.String("pokemon-id", pkm.Id),
	}

	r.logger.Info("creating pokemon", lf...)

	_, err := r.coll.InsertOne(context.Background(), r.initializeDtoFromModel(pkm))

	if err != nil {
		lf = append(lf, zap.NamedError("error", err))
		r.logger.Error("unexpected error at creating pokemon", lf...)
	}

	return err
}

func (r *PokemonMongo) CreateManyPokemon(pkms []*domain.Pokemon) error {

	pkmsd := make([]interface{}, len(pkms))

	for _, pkm := range pkms {
		pkmsd = append(pkmsd, r.initializeDtoFromModel(pkm))
	}

	_, err := r.coll.InsertMany(context.Background(), pkmsd)

	if err != nil {
		lf := []zap.Field{
			zap.String("logger", "PokemonMongo"),
			zap.String("sub-logger", "CreateManyPokemon"),
		}

		lf = append(lf, zap.NamedError("error", err))
		r.logger.Error("unexpected error at creating many pokemon", lf...)
	}

	return err
}

func NewPokemonRepository(cm *mongo.Database, c *config.Config, logger *zap.Logger) domain.PokemonRepository {

	return &PokemonMongo{
		coll:   cm.Collection(c.PokemonMongoCollection),
		logger: logger,
	}
}
