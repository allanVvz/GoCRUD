package routes

import (
	"github.com/allanVvz/GoCRUD/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {

	r.GET("/getDriver/:userId", controller.FindDriverById)
	r.GET("/getDriver/:DriverReg", controller.FindDriverByReg)
	r.GET("/getDriver/:DriverRg", controller.FindDriverByRg)
	r.GET("/getSalary/:DriverId", controller.GetSalary)
	r.POST("/createDriver", controller.CreateDriver)
	r.POST("/Journey/", controller.CreateJourney)
	r.PUT("/updateDriver/:driverId", controller.UpdateDriver)
	r.DELETE("/deleteDriver/:driverId", controller.FireDriver)

}
