package routes

import (
	"github.com/Johnman67112/game_management_api/controllers"
	"github.com/Johnman67112/game_management_api/docs"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Handlers() {
	//Set router
	r := gin.Default()

	//Swag
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//Games
	r.GET("/games/:name/:plataform", controllers.GetGame)

	r.GET("/games", controllers.GetGames)

	r.POST("/games", controllers.AddGame)

	r.PATCH("/games", controllers.UpdateGameState)

	r.DELETE("/games/:name/:plataform", controllers.RemoveGame)

	//Run
	r.Run()
}
