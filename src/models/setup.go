package models

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func getEnvVariables(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

// ConnectDatabase function stablish a connection to MongoDB and it colletion
func ConnectDatabase() (ctx context.Context, collection *mongo.Collection, cancel context.CancelFunc, err error) {
	var clientOptions *options.ClientOptions

	localDB := "mongodb://localhost:27017"
	remoteDB := "mongodb+srv://" + getEnvVariables("REMOTE_DB_USER") + ":" + getEnvVariables("REMOTE_DB_PW") + "@montertest.it8hx.mongodb.net/" + getEnvVariables("DB") + "?retryWrites=true&w=majority"

	db := getEnvVariables("DB")
	dbCollection := getEnvVariables("COLLECTION")

	if getEnvVariables("APP_ENV") == "production" {
		clientOptions = options.Client().ApplyURI(remoteDB)
	} else if getEnvVariables("APP_ENV") == "development" {
		clientOptions = options.Client().ApplyURI(localDB)
	}

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
