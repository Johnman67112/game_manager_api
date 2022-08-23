package main

import (
	"github.com/Johnman67112/game_management_api/database"
	"github.com/Johnman67112/game_management_api/routes"
	"github.com/joho/godotenv"
)

//Version: 0.1

func main() {
	godotenv.Load(".env")
	database.Connect()
	routes.Handlers()
}
