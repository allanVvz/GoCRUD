package main

import (
	"fmt"
	"log"
	"os"

	"github.com/allanVvz/GoCRUD/src/config/logger"
	"github.com/allanVvz/GoCRUD/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("Starting the application")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")

	}
	
	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup)

	if err := router.Run(":8080"); err != nil {
		fmt.Println("Error starting server")
		os.Exit(1)
	}
}
