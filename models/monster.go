package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Monster definition
type Monster struct {
	ID             primitive.ObjectID `bson: "_id"				json: "_id"`
	CreatedAt      time.Time          `bson: "createdAt"		json: "createdAt"`
	Name           string             `bson: "name"		 		json: "name"`
	Moves          []string           `bson: "moves"			json: "moves"`
	Instinct       string             `bson: "instinct"	 		json: "instinct"`
	Description    string             `bson: "description"		json: "description"`
	Attack         string             `bson: "attacker"			json: "attack"`
	AttackTags     []string           `bson: "attack_tags"		json: "attack_tags"`
	Damage         string             `bson: "damage"			json: "damage"`
	MonsterTags    []string           `bson: "monster_tags" 	json: "monster_tags"`
	HP             int32              `bson: "hp"				json: "hp"`
	Armor          int32              `bson: "armor"			json: "armor"`
	SpecialQuality []string           `bson: "special_quality" 	json: "special_quality"`
	Setting        string             `bson: "setting"	 		json: "setting"`
	Source         string             `bson: "source" 			json: "source"`
}
