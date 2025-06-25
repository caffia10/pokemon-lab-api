package pokemontype

import "go.mongodb.org/mongo-driver/bson/primitive"

type pokemonTypeDto struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type pokemonTypeMongoDto struct {
	Id   primitive.ObjectID `bson:"_id"`
	Name string             `bson:"name"`
}

type pokemonTypeScyllasDto struct {
	Id   string `db:"id"`
	Name string `db:"name"`
}
