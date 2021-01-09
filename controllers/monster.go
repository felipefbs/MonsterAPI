package controllers

import (
	"net/http"

	"github.com/felipefbs/MonsterAPI/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetAllMonsters function returns every monsters in database
func GetAllMonsters(c *gin.Context) {
	monsterCtx, monsterCollection := models.ConnectDatabase()

	var monsters []*models.Monster

	cur, err := monsterCollection.Find(monsterCtx, bson.D{{}})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	for cur.Next(monsterCtx) {
		var m models.Monster
		err := cur.Decode(&m)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		monsters = append(monsters, &m)
	}

	if err := cur.Err(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	cur.Close(monsterCtx)

	if len(monsters) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": mongo.ErrNoDocuments})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": monsters})
}

// CreateMonster function
func CreateMonster(c *gin.Context) {
	monsterCtx, monsterCollection := models.ConnectDatabase()
	var input models.Monster
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	monster := models.Monster{Name: input.Name, Area: input.Area}

	_, err := monsterCollection.InsertOne(monsterCtx, monster)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": monster})
}
