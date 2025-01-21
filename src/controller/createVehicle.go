package controller

import (
	"fmt"

	"github.com/allanVvz/GoCRUD/src/config/rest_err"
	"github.com/allanVvz/GoCRUD/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateVehicle(c *gin.Context) {

	var vehicleRequest request.VehicleRequest

	if err := c.ShouldBindJSON(&vehicleRequest); err != nil {
		restErr := rest_err.NewBadRequestError(
			fmt.Sprintf("Invalid JSON body, there are some incorrect fields: %s ", err.Error()))
		c.JSON(restErr.Code, restErr)
		return
	}
	fmt.Println(vehicleRequest)
}
