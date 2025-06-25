package server

import (
	"context"
	"pokemon-lab-api/internal/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDatabase(c *config.Config) *mongo.Database {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(c.MongoUri))

	if err != nil {
		panic(err)
	}

	return client.Database(c.MongoDatabase)
}
