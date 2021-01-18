package main

import (
	"net/http"

	"github.com/felipefbs/MonsterAPI/controllers"
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
)

func main() {
	Router := gin.Default()

	Router.Use(favicon.New("./assets/favicon.ico"))

	api := Router.Group("/api/v1")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "Hello, Monster!"})
		})
		api.GET("/monsters", controllers.GetAllMonsters)
		api.GET("/monsters/:setting", controllers.GetMonstersBySetting)
		api.POST("/monsters", controllers.CreateMonster)
	}

	Router.Run()

}
