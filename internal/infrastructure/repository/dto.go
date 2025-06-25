package repository

import "go.mongodb.org/mongo-driver/bson/primitive"

type evolutionDto struct {
	Id          string `db:"id"`
	Requirement string `db:"requirement"`
}

type pokemonDto struct {
	Id         string         `db:"id"`
	Name       string         `db:"name"`
	Category   string         `db:"category"`
	Weight     string         `db:"weight"`
	ImgUrl     string         `db:"img_url"`
	Types      []string       `db:"types"` // save only the ids
	Evolutions []evolutionDto `db:"evolutions"`
}

type pokemonTypeMongoDto struct {
	Id   primitive.ObjectID `bson:"_id"`
	Name string             `bson:"name"`
}

type pokemonTypeDto struct {
	Id   string `db:"id"`
	Name string `db:"name"`
}
