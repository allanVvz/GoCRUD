package routes

import (
	"github.com/allanVvz/GoCRUD/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine, driverController *controller.DriverController) {
	// Grupo de rotas para motoristas
	drivers := router.Group("/drivers")
	{
		drivers.GET("/", driverController.FindDriver) // Usa query parameters para buscar por ID, REG ou RG
		drivers.POST("/", driverController.CreateDriver)
		drivers.PUT("/:driverId", driverController.UpdateDriver)
	}

	// Grupo de rotas para veículos
	vehicles := router.Group("/vehicles")
	{
		//vehicles.GET("/", controller.FindVehicle) // Usa query parameters para buscar veículos por ID, Modelo, Placa ou Ano
		vehicles.POST("/", controller.CreateVehicle)
		vehicles.PUT("/:vehicleId", controller.UpdateVehicle)
		vehicles.DELETE("/:vehicleId", controller.FireVehicle)
	}

	// Grupo de rotas para jornadas
	journeys := router.Group("/journeys")
	{
		journeys.GET("/", controller.ListJourneys) // Usa query parameters para diferenciar tipos de jornada
		journeys.POST("/", controller.CreateJourney)
	}
}
