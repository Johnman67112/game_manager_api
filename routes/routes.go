package routes

import "github.com/gin-gonic/gin"

func Handlers() {
	//Set router
	r := gin.Default()

	r.Group("/api/v1")

	r.GET("/game")
	r.POST("/game")
	r.DELETE("/game")

	//Run
	r.Run()
}
