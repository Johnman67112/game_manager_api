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

	database.DB.Where(&models.Game{Name: name, Plataform: plataform}).First(&game)

	if game.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Game not found"})
		return
	}

	c.JSON(http.StatusOK, game)
}

func GetGames(c *gin.Context) {
	var games []models.Game
	database.DB.Find(&games)
	c.JSON(200, games)
}

func AddGame(c *gin.Context) {
	var game, gameDB models.Game

	if err := c.ShouldBindJSON(&game); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Where(&models.Game{Name: game.Name, Plataform: game.Plataform}).First(&gameDB)

	if game.Name == gameDB.Name && game.Plataform == gameDB.Plataform {
		c.JSON(http.StatusConflict, gin.H{
			"Duplicated": "Game already exists"})
		return
	}

	database.DB.Create(&game)

	c.JSON(http.StatusOK, "Game Created Sucessfully")
}

func UpdateGameState(c *gin.Context) {
	var game, gameDB models.Game

	if err := c.ShouldBindJSON(&game); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Where(&models.Game{Name: game.Name, Plataform: game.Plataform}).First(&gameDB)

	if gameDB.ID == 0 {
		c.JSON(http.StatusConflict, gin.H{
			"Not found": "Game not found"})
		return
	}

	database.DB.Model(&gameDB).UpdateColumns(&game)

	c.JSON(http.StatusOK, "Game State Updated Sucessfully")
}

func RemoveGame(c *gin.Context) {
	var game models.Game

	name := c.Param("name")
	plataform := c.Param("plataform")

	database.DB.Where(&models.Game{Name: name, Plataform: plataform}).First(&game)

	if game.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Game not found"})
		return
	}

	database.DB.Delete(&game, game.ID)

	c.JSON(http.StatusOK, "Game Deleted Sucessfully")
}
