package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Monster definition
type Monster struct {
	ID   primitive.ObjectID `bson:_id`
	Name string             `bson:name`
	Area string             `bson:area`
}
