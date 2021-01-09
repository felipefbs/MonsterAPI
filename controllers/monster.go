package controllers

import (
	"net/http"

	"github.com/felipefbs/MonsterAPI/models"
	"github.com/gin-gonic/gin"
)

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
