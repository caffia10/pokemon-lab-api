package mongorepo

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type pokemonDto struct {
	Id       primitive.ObjectID    `bson:"_id"`
	Name     string                `bson:"name"`
	Category string                `bson:"category"`
	Weight   string                `bson:"weight"`
	ImgUrl   string                `bson:"img_url"`
	Types    []*primitive.ObjectID `bson:"type_ids"`
}
