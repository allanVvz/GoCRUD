package main

import (
	"context"
	"log"
	"time"

	"github.com/allanVvz/GoCRUD/src/controller"
	"github.com/allanVvz/GoCRUD/src/controller/routes"
	"github.com/allanVvz/GoCRUD/src/model/repository"
	"github.com/allanVvz/GoCRUD/src/model/service"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Conectar ao MongoDB
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalf("Failed to create MongoDB client: %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}
	db := client.Database("gocrud")

	// Inicializar o repositório, serviço e controlador
	driverRepository := repository.NewDriverRepository(db)
	driverService := service.NewDriverService(driverRepository)
	driverController := controller.NewDriverController(driverService)

	// Inicializar o roteador
	router := gin.Default()
	routes.InitRoutes(router, driverController)

	// Iniciar o servidor
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
