package controllers

import (
	"net/http"

	"github.com/Johnman67112/game_management_api/database"
	"github.com/Johnman67112/game_management_api/models"
	"github.com/gin-gonic/gin"

	_ "github.com/swaggo/swag/example/celler/httputil"
	_ "github.com/swaggo/swag/example/celler/model"
)

//Get Game godoc
// @Summary      Show a game
// @Description  Route to show a game with params name and plataform
// @Tags         Games
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.Game
// @Failure      404  {object}  httputil.HTTPError
// @Router       /games/{name}/{plataform} [get]
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

//Get Games godoc
// @Summary      Show all games
// @Description  Route to show all games
// @Tags         Games
// @Accept       json
// @Produce      json
// @Success      200  {object}  []models.Game
// @Failure      400  {object}  httputil.HTTPError
// @Router       /games [get]
func GetGames(c *gin.Context) {
	var games []models.Game
	database.DB.Find(&games)
	c.JSON(200, games)
}

//Add Game godoc
// @Summary      Adds a new game on database
// @Description  With params Name, Plataform, Status, Percentage adds a new game
// @Tags         Games
// @Accept       json
// @Produce      json
// @Param        game  body  models.Game  true  "Game Model"
// @Success      200  {object}  string
// @Failure      400  {object}  httputil.HTTPError
// @Failure      409  {object}  httputil.HTTPError
// @Router       /games [post]
func AddGame(c *gin.Context) {
	var game, gameDB models.Game

	if err := c.ShouldBindJSON(&game); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	if err := models.GameValidator(&game); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error()})
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

//Update Game State godoc
// @Summary      Update game status and parcentage
// @Description  With params Name Plataform, Status and Percentage updates game state
// @Tags         Games
// @Accept       json
// @Produce      json
// @Param        game  body  models.Game  true  "Game Model"
// @Success      200  {object}  string
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Router       /games [patch]
func UpdateGameState(c *gin.Context) {
	var game, gameDB models.Game

	if err := c.ShouldBindJSON(&game); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	if err := models.GameValidator(&game); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error()})
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

//Remove Game godoc
// @Summary      Removes a new game on database
// @Description  With params Name, Plataform, Status, Percentage removes a new game
// @Tags         Games
// @Accept       json
// @Produce      json
// @Success      200  {object}  string
// @Failure      404  {object}  httputil.HTTPError
// @Router       /games/{name}/{plataform} [delete]
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
