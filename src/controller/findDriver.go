package controller

import (
	"net/http"

	"github.com/allanVvz/GoCRUD/src/config/logger"
	"github.com/allanVvz/GoCRUD/src/config/rest_err"
	"github.com/allanVvz/GoCRUD/src/controller/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// FindDriver busca um motorista pelo ID, Registro ou RG usando query parameters.
func (dc *DriverController) FindDriver(c *gin.Context) {
	driverID := c.Query("driverId")
	driverReg := c.Query("driverReg")
	driverRg := c.Query("driverRg")
	getSalary := c.Query("getSalary") // Flag opcional para retornar apenas o salário

	if driverID == "" && driverReg == "" && driverRg == "" {
		logger.Error("No valid query parameter provided in FindDriver",
			nil, zap.String("journey", "findDriver"))
		c.JSON(http.StatusBadRequest, rest_err.NewBadRequestError("You must provide driverId, driverReg, or driverRg"))
		return
	}

	var driver response.DriverResponse
	var err *rest_err.RestErr

	// Decide qual método chamar com base no parâmetro recebido
	if driverID != "" {
		driver, err = dc.Service.FindDriverByID(driverID)
	} else if driverReg != "" {
		driver, err = dc.Service.FindDriverByReg(driverReg)
	} else if driverRg != "" {
		driver, err = dc.Service.FindDriverByRg(driverRg)
	}

	if err != nil {
		logger.Error("Error finding driver", err,
			zap.String("journey", "findDriver"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("Driver found successfully", zap.String("journey", "findDriver"))

	// Se a query "getSalary" for true, retorna apenas o salário
	if getSalary == "true" {
		c.JSON(http.StatusOK, gin.H{
			"driverId": driver.Id,
			"salary":   driver.Salary,
		})
		return
	}

	// Retorna os dados completos do motorista
	c.JSON(http.StatusOK, driver)
}
