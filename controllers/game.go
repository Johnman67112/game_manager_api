package controllers

import "github.com/gin-gonic/gin"

func GetGameRequest(c *gin.Context) {
	name := c.Param("name")
	plataform := c.Param("plataform")
}
