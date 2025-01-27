package controller

import (
	"net/http"

	"github.com/allanVvz/GoCRUD/src/config/logger"
	"github.com/allanVvz/GoCRUD/src/config/rest_err"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// FindDriverById é uma função que recebe um driverId e retorna os dados do motorista
func (dc *DriverController) FindDriverById(c *gin.Context) {
	driverID := c.Param("driverId")
	if driverID == "" {
		logger.Error("Invalid driverId received in FindDriverById", nil,
			zap.String("journey", "findDriverById"),
			zap.String("driverID", driverID),
		)
		c.JSON(http.StatusBadRequest, rest_err.NewBadRequestError("Driver ID is required"))
		return
	}

	// Chama o serviço para buscar o motorista pelo ID
	driver, err := dc.Service.FindDriverByID(driverID)
	if err != nil {
		logger.Error("Error finding driver in FindDriverById", err,
			zap.String("journey", "findDriverById"),
			zap.String("driverID", driverID),
		)
		c.JSON(err.Code, err)
		return
	}

	logger.Info("Driver found successfully in FindDriverById",
		zap.String("journey", "findDriverById"),
		zap.String("driverID", driverID),
	)
	c.JSON(http.StatusOK, driver)
}


func FindDriverByReg(c *gin.Context) {

}

func FindDriverByRg(c *gin.Context) {

}

func GetSalary(c *gin.Context) {

}
