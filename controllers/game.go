package controllers

import (
	"net/http"

	"github.com/Johnman67112/game_management_api/database"
	"github.com/Johnman67112/game_management_api/models"
	"github.com/gin-gonic/gin"
)

func GetGame(c *gin.Context) {
	var game models.Game
	name := c.Param("name")
	plataform := c.Param("plataform")

	database.DB.First(&game, name, plataform)

	if game.Name == "" && game.Plataform == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Game not found"})
		return
	}

	c.JSON(http.StatusOK, game)
}
