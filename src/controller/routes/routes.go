package routes

import (
	"github.com/allanVvz/GoCRUD/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine, driverController *controller.DriverController) {
	// Grupo de rotas para motoristas
	drivers := router.Group("/drivers")
	{
		drivers.GET("/id/:driverId", driverController.FindDriverById)
		drivers.GET("/reg/:driverReg", controller.FindDriverByReg)
		drivers.GET("/rg/:driverRg", controller.FindDriverByRg)
		drivers.GET("/salary/:driverId", controller.GetSalary)
		drivers.POST("/", driverController.CreateDriver)
		drivers.PUT("/:driverId", driverController.UpdateDriver)
		//drivers.DELETE("/:driverId", driverController.FireDriver)
	}

	// Grupo de rotas para veículos
	vehicles := router.Group("/vehicles")
	{
		vehicles.GET("/id/:vehicleId", controller.FindVehicleById)
		vehicles.GET("/model/:vehicleModel", controller.FindVehicleByModel)
		vehicles.GET("/plate/:vehiclePlate", controller.FindVehicleByPlate)
		vehicles.GET("/date/:vehicleYear", controller.FindVehicleByPurchaseDate)
		vehicles.POST("/", controller.CreateVehicle)
		vehicles.PUT("/:vehicleId", controller.UpdateVehicle)
		vehicles.DELETE("/:vehicleId", controller.FireVehicle)
	}

	// Grupo de rotas para jornadas
	journeys := router.Group("/journeys")
	{
		journeys.GET("/", controller.List2wJourneys)
		journeys.GET("/2w", controller.ListJourneys)
		journeys.POST("/", controller.CreateJourney)
	}
}
