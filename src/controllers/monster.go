package controllers

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/felipefbs/MonsterAPI/src/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetAllMonsters function returns every monsters in database
func GetAllMonsters(c *gin.Context) {
	filter := bson.M{}

	a := c.Query("attack_tags")
	m := c.Query("monster_tags")
	n := c.Query("name")

	if a != "" {
		filter["attack_tags"] = bson.M{"$in": toLowerCase(strings.Split(a, ","))}
	}
	if m != "" {
		filter["monster_tags"] = bson.M{"$in": toLowerCase(strings.Split(m, ","))}
	}
	if n != "" {
		filter["name"] = strings.ToLower(n)
	}

	monsters, err := filterMonsters(filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"numberOfItens": len(monsters), "data": monsters})
}

// GetMonstersBySetting returns all monsters from a specific setting
func GetMonstersBySetting(c *gin.Context) {
	s := c.Param("setting")
	a := c.Query("attack_tags")
	m := c.Query("monster_tags")

	filter := bson.M{
		"setting": s,
	}

	if a != "" {
		filter["attack_tags"] = bson.M{"$in": toLowerCase(strings.Split(a, ","))}
	}
	if m != "" {
		filter["monster_tags"] = bson.M{"$in": toLowerCase(strings.Split(m, ","))}
	}

	monsters, err := filterMonsters(filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"numberOfItens": len(monsters), "data": monsters})
}

// GetMonstersByName function handler returns a monster based on it name
func GetMonstersByName(c *gin.Context) {
	n := c.Param("name")
	name := strings.ToLower(n)

	filter := bson.M{
		"name": name,
	}

	monsters, err := filterMonsters(filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": monsters[0]})
}

// filterMonster function returns to client a monster slice based on a filter
func filterMonsters(filter interface{}) (monsters []models.Monster, err error) {
	monsterCtx, monsterCollection, cancel, err := models.ConnectDatabase()
	defer cancel()
	if err != nil {
		return
	}

	cur, err := monsterCollection.Find(monsterCtx, filter)
	if err != nil {
		return
	}

	for cur.Next(monsterCtx) {
		var m models.Monster
		err = cur.Decode(&m)
		if err != nil {
			return
		}
		monsters = append(monsters, m)
	}

	if err = cur.Err(); err != nil {
		return
	}

	cur.Close(monsterCtx)

	if len(monsters) == 0 {
		err = mongo.ErrNoDocuments
		return
	}

	return monsters, nil
}

// CreateMonster function
func CreateMonster(c *gin.Context) {
	monsterCtx, monsterCollection, cancel, err := models.ConnectDatabase()
	defer cancel()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

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
		Name:             strings.ToLower(strings.TrimSpace(monster.Name)),
		Moves:            toLowerCase(monster.Moves),
		Instinct:         strings.TrimSpace(monster.Instinct),
		Description:      strings.TrimSpace(monster.Description),
		Attack:           strings.TrimSpace(monster.Attack),
		AttackTags:       toLowerCase(monster.AttackTags),
		MonsterTags:      toLowerCase(monster.MonsterTags),
		Damage:           strings.TrimSpace(monster.Damage),
		HP:               monster.HP,
		Armor:            monster.Armor,
		SpecialQualities: toLowerCase(monster.SpecialQualities),
		Setting:          strings.TrimSpace(monster.Setting),
		Source:           strings.TrimSpace(monster.Source),
	}

	_, err := collection.InsertOne(ctx, monster)

	return monster, err
}

func toLowerCase(s []string) []string {
	var aux []string
	for _, v := range s {
		aux = append(aux, strings.ToLower(strings.TrimSpace(v)))
	}
	return aux
}

// UpdateMonster function handler changes some monster properties based on it name
func UpdateMonster(c *gin.Context) {
	var input models.Monster
	n := strings.ToLower(c.Param("name"))

	monsterCtx, monsterCollection, cancel, err := models.ConnectDatabase()
	defer cancel()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	monster, err := filterMonsters(bson.M{"name": n})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if monster[0].Source == "core rulebook" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You cannot change Monsters from core rulebook"})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	monsterUpdate := bson.M{
		"name":              input.Name,
		"moves":             input.Moves,
		"instinct":          input.Instinct,
		"description":       input.Description,
		"attack":            input.Attack,
		"attack_tags":       input.AttackTags,
		"damage":            input.Damage,
		"monster_tags":      input.MonsterTags,
		"hp":                input.HP,
		"armor":             input.Armor,
		"special_qualities": input.SpecialQualities,
		"setting":           input.Setting,
		"source":            input.Source,
	}

	update := bson.M{"$set": monsterUpdate}

	m := monsterCollection.FindOneAndUpdate(
		monsterCtx,
		bson.M{"name": n},
		update,
	)
	if m.Err() != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": m.Err().Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": monsterUpdate})

}
