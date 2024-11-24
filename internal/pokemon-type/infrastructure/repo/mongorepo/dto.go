package mongorepo

import "go.mongodb.org/mongo-driver/bson/primitive"

type pokemonTypeDto struct {
	Id   primitive.ObjectID `bson:"_id"`
	Name string             `bson:"name"`
}
