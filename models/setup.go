package models

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
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
func ConnectDatabase(c *gin.Context) (ctx context.Context, collection *mongo.Collection, cancel context.CancelFunc) {
	clientOptions := options.Client().ApplyURI(dbURL)

	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err = client.Ping(ctx, readpref.Primary())

	collection = client.Database(db).Collection(dbCollection)

	return
}
