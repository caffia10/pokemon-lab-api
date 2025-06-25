package repository

import (
	"context"
	"errors"
	"log"
	"pokemon-lab-api/internal/domain"
	"pokemon-lab-api/internal/infrastructure/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

var (
	ErrInvalidIds        = errors.New("all ids to find the types are invalids")
	ErrUnexpectedOverall = errors.New("unexpected error at decoding all types")
)

type PokemonTypeMongoRepository struct {
	coll   *mongo.Collection
	logger *zap.Logger
}

func (r *PokemonTypeMongoRepository) RetriveById(id string) (*domain.PokemonType, error) {

	idpo, errIdpo := primitive.ObjectIDFromHex(id)

	lf := []zap.Field{
		zap.String("logger", "PokemonTypeMongoRepository"),
		zap.String("sub-logger", "RetriveById"),
		zap.String("pokemon-id", id),
	}

	r.logger.Info("retrieving pokemon", lf...)

	if errIdpo != nil {
		log.Printf("[PokemonTypeMongoRepository][RetriveById] fail to create the primitive objectId %s\n", id)
		return nil, errIdpo
	}

	pkmt := new(pokemonTypeDto)
	err := r.coll.FindOne(context.TODO(), bson.D{{Key: "_id", Value: idpo}}).
		Decode(pkmt)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			r.logger.Warn("no document was found")
			return nil, domain.ErrNotFoundPokemonType
		}
		lf = append(lf, zap.NamedError("error", err))
		r.logger.Error("unexpected error at retriving pokemon", lf...)
		return nil, err
	}

	return &domain.PokemonType{
		Id:   pkmt.Id,
		Name: pkmt.Name,
	}, nil
}

func (r *PokemonTypeMongoRepository) RetriveAll() ([]domain.PokemonType, error) {

	filter := bson.M{}
	ctx := context.TODO()
	cursor, err := r.coll.Find(ctx, filter)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Printf("[PokemonTypeMongoRepository][RetriveByIds]no document was found for the id\n")
			return nil, domain.ErrNotFoundPokemonType
		}
		log.Printf("[PokemonTypeMongoRepository][RetriveByIds]unexpected error at retriving pokemon types by Ids, detail:%s\n", err.Error())
		return nil, err
	}

	defer cursor.Close(ctx)
	var pkmts []domain.PokemonType

	for cursor.Next(ctx) {
		var pkmt domain.PokemonType
		errDecode := cursor.Decode(&pkmt)

		if errDecode != nil {
			log.Printf("[PokemonTypeMongoRepository][RetriveByIds] fail to decode the pokemon type %s\n", errDecode.Error())
			continue
		}

		pkmts = append(pkmts, pkmt)
	}

	if len(pkmts) == 0 {
		log.Printf("[PokemonTypeMongoRepository][RetriveByIds] %s\n", ErrUnexpectedOverall.Error())
		return nil, ErrUnexpectedOverall
	}

	return pkmts, nil
}

func (r *PokemonTypeMongoRepository) Create(pkmt *domain.PokemonType) error {

	log.Printf("[PokemonTypeMongoRepository][CreatePokemonType] creating pokemon type %s \n", pkmt.Name)

	pkmd := &pokemonTypeDto{
		Id:   primitive.NewObjectID().String(),
		Name: pkmt.Name,
	}
	_, err := r.coll.InsertOne(context.Background(), pkmd)

	if err != nil {
		log.Printf("[PokemonTypeMongoRepository][CreatePokemonType]unexpected error at creating pokemon type %v, details: %s \n", pkmt, err.Error())
	}

	return err
}

func NewPokemonTypeMongoRepository(cm *mongo.Database, c *config.Config) domain.PokemonTypeRepository {

	return &PokemonTypeMongoRepository{
		coll: cm.Collection(c.PokemonTypeMongoCollection),
	}
}
