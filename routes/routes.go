package routes

import (
	"github.com/Johnman67112/game_management_api/controllers"
	"github.com/gin-gonic/gin"
)

func Handlers() {
	//Set router
	r := gin.Default()

	//Games
	r.GET("/games/:name/:plataform", controllers.GetGame)

	r.GET("/games", controllers.GetGames)

	r.POST("/games", controllers.AddGame)

	r.DELETE("/games/:name/:plataform", controllers.RemoveGame)

	//Run
	r.Run()
}
