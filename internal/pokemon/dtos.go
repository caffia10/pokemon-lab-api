package pokemon

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type pokemonMongoDto struct {
	Id       primitive.ObjectID    `bson:"_id"`
	Name     string                `bson:"name"`
	Category string                `bson:"category"`
	Weight   string                `bson:"weight"`
	ImgUrl   string                `bson:"img_url"`
	Types    []*primitive.ObjectID `bson:"type_ids"`
}

type pokemonTypeDto struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type pokemonRequestDto struct {
	Name     string            `json:"name"`
	Weight   string            `json:"weight"`
	Category string            `json:"category"`
	ImgUrl   string            `json:"img_url"`
	Types    []*pokemonTypeDto `json:"types"`
}

type pokemonScyllaDto struct {
	Id       string    `db:"id"`
	Name     string    `db:"name"`
	Category string    `db:"category"`
	Weight   string    `db:"weight"`
	ImgUrl   string    `db:"img_url"`
	Types    []*string `db:"types"` // save only the id
}
