package main

import (
	"github.com/felipefbs/MonsterAPI/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	Router := gin.Default()

	api := Router.Group("/api/v1")
	{
		api.GET("/", controllers.GetAllMonsters)
		api.POST("/", controllers.CreateMonster)
	}

	Router.Run()

}
