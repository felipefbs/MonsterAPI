package main

import (
	"net/http"

	"github.com/felipefbs/MonsterAPI/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	Router := gin.Default()

	api := Router.Group("/api/v1")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "Hello, Monster!"})
		})
		api.GET("/monsters", controllers.GetAllMonsters)
		api.POST("/monsters", controllers.CreateMonster)
	}

	Router.Run()

}
