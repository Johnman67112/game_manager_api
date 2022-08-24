package routes

import (
	"github.com/Johnman67112/game_management_api/controllers"
	"github.com/gin-gonic/gin"
)

func Handlers() {
	//Set router
	r := gin.Default()

	r.Group("/api/v1")

	r.GET("/games/:name/:plataform", controllers.GetGame)
	r.POST("/games", controllers.AddGame)
	r.DELETE("/games")

	//Run
	r.Run()
}
