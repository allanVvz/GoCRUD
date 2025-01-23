package routes

import (
	"github.com/allanVvz/GoCRUD/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {

	r.GET("/getDriver/Id/:DriverId", controller.FindDriverById)
	r.GET("/getDriver/Reg/:DriverReg", controller.FindDriverByReg)
	r.GET("/getDriver/Rg/:DriverRg", controller.FindDriverByRg)
	r.GET("/getSalary/:DriverId", controller.GetSalary)
	r.GET("/getVehicle/Id/:vehicleId", controller.FindVehicleById)
	r.GET("/getVehicle/Model/:vehicleModel", controller.FindVehicleByModel)
	r.GET("/getVehicle/Plate/:vehiclePlate", controller.FindVehicleByPlate)
	r.GET("/getVehicle/Date/:vehicleYear", controller.FindVehicleByPurchaseDate)
	r.GET("/getJourney/", controller.List2wJourneys)
	r.GET("/getJourney/2w/", controller.ListJourneys)

	r.POST("/Driver", controller.createDriver)
	r.POST("/Vehicles/", controller.CreateVehicle)
	r.POST("/Journey/", controller.CreateJourney)

	r.PUT("/updateDriver/:driverId", controller.UpdateDriver)
	r.PUT("/updateVehicle/:driverId", controller.UpdateVehicle)
	r.PUT("/deleteDriver/:driverId", controller.FireDriver)
	r.PUT("/deleteVehicle/:vehicleId", controller.FireVehicle)

}
