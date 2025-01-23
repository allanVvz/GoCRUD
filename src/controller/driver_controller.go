package controller

import (
	"net/http"

	"github.com/allanVvz/GoCRUD/src/controller/model/request"
	"github.com/allanVvz/GoCRUD/src/model/service"
	"github.com/gin-gonic/gin"
)

type DriverController struct {
	Service *service.DriverService
}

func NewDriverController(service *service.DriverService) *DriverController {
	return &DriverController{
		Service: service,
	}
}

func (dc *DriverController) createDriver(c *gin.Context) {
	var driverRequest request.DriverRequest

	if err := c.ShouldBindJSON(&driverRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	driverResponse, err := dc.Service.CreateDriver(driverRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, driverResponse)
}
