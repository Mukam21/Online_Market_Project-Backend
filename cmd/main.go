package main

import (
	"log"
	"online_market_project/pkg/config"
	"online_market_project/pkg/database"
	"online_market_project/pkg/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()

	database.ConnectDatabase()

	router := gin.Default()

	routes.InitializeRoutes(router)

	if err := router.Run(config.Config.ServerPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
