package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	dbURL        = "mongodb://localhost:27017"
	db           = "MonsterDB"
	dbCollection = "monsters"
)

// ConnectDatabase function stablish a connection to MongoDB and it colletion
func ConnectDatabase() (ctx context.Context, collection *mongo.Collection, cancel context.CancelFunc, err error) {
	clientOptions := options.Client().ApplyURI(dbURL)

	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return
	}

	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		return
	}

	err = client.Ping(ctx, readpref.Primary())

	collection = client.Database(db).Collection(dbCollection)

	return ctx, collection, cancel, nil
}
