package routes

import (
	"github.com/allanVvz/GoCRUD/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {

	r.GET("/getUserById/:userId", controller.FindDriverById)
	r.GET("/getDriverByEmail/:DriverEmail", controller.FindDriverByReg)
	r.POST("/createDriver", controller.CreateDriver)
	r.PUT("/updateDriver/:driverId", controller.UpdateDriver)
	r.DELETE("/deleteDriver/:driverId", controller.DeleteDriver)

}
