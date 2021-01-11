package controllers

import (
	"context"
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

	var input []models.Monster
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var mSlice []models.Monster
	for _, monster := range input {
		m, err := insertMonster(monsterCtx, monsterCollection, monster)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		mSlice = append(mSlice, m)

	}

	c.JSON(http.StatusOK, gin.H{"data": mSlice})
}

func insertMonster(ctx context.Context, collection *mongo.Collection, monster models.Monster) (models.Monster, error) {
	monster = models.Monster{
		ID:               primitive.NewObjectID(),
		CreatedAt:        time.Now(),
		Name:             monster.Name,
		Moves:            monster.Moves,
		Instinct:         monster.Instinct,
		Description:      monster.Description,
		Attack:           monster.Attack,
		AttackTags:       monster.AttackTags,
		MonsterTags:      monster.MonsterTags,
		Damage:           monster.Damage,
		HP:               monster.HP,
		Armor:            monster.Armor,
		SpecialQualities: monster.SpecialQualities,
		Setting:          monster.Setting,
		Source:           monster.Source,
	}

	_, err := collection.InsertOne(ctx, monster)

	return monster, err
}
