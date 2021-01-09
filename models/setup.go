package models

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectDatabase function stablish a connection to MongoDB and it colletion
func ConnectDatabase(c *gin.Context) (ctx context.Context, collection *mongo.Collection) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Fatal(err)
	}

	collection = client.Database("MonsterDB").Collection("monsters")

	return ctx, collection
}
