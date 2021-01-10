package controllers

import (
	"net/http"
	"time"

	"github.com/felipefbs/MonsterAPI/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetAllMonsters function returns every monsters in database
func GetAllMonsters(c *gin.Context) {
	filter := bson.D{{}}

	filterMonsters(filter, c)
}

// filterMonster function returns to client a monster slice based on a filter
func filterMonsters(filter interface{}, c *gin.Context) {
	monsterCtx, monsterCollection, cancel := models.ConnectDatabase(c)
	defer cancel()

	var monsters []*models.Monster

	cur, err := monsterCollection.Find(monsterCtx, filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for cur.Next(monsterCtx) {
		var m models.Monster
		err := cur.Decode(&m)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		monsters = append(monsters, &m)
	}

	if err := cur.Err(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cur.Close(monsterCtx)

	if len(monsters) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": mongo.ErrNoDocuments.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": monsters})
}

// CreateMonster function
func CreateMonster(c *gin.Context) {
	monsterCtx, monsterCollection, cancel := models.ConnectDatabase(c)
	defer cancel()

	var input models.Monster
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	monster := models.Monster{
		ID:             primitive.NewObjectID(),
		CreatedAt:      time.Now(),
		Name:           input.Name,
		Moves:          input.Moves,
		Instinct:       input.Instinct,
		Description:    input.Description,
		Attack:         input.Attack,
		AttackTags:     input.AttackTags,
		MonsterTags:    input.MonsterTags,
		Damage:         input.Damage,
		HP:             input.HP,
		Armor:          input.Armor,
		SpecialQuality: input.SpecialQuality,
		Setting:        input.Setting,
		Source:         input.Source,
	}

	_, err := monsterCollection.InsertOne(monsterCtx, monster)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": monster})
}
