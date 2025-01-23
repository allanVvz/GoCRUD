package controller

import (
	"net/http"

	"github.com/allanVvz/GoCRUD/src/config/logger"
	"github.com/allanVvz/GoCRUD/src/config/validation"
	"github.com/allanVvz/GoCRUD/src/controller/model/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (dc *DriverController) CreateDriver(c *gin.Context) {
	logger.Info("Init CreateDriver controller",
		zap.String("journey", "createDriver"),
	)

	// Bind JSON diretamente no modelo request.DriverRequest
	var driverRequest request.DriverRequest
	if err := c.ShouldBindJSON(&driverRequest); err != nil {
		logger.Error("Invalid JSON received in CreateDriver", err,
			zap.String("journey", "createDriver"))
		validationError := validation.ValidateDriverError(err)
		c.JSON(validationError.Code, validationError)
		return
	}

	// Chama o serviço para criar o motorista
	driverResponse, err := dc.Service.CreateDriver(driverRequest)
	if err != nil {
		logger.Error("Error calling CreateDriver service", err,
			zap.String("journey", "createDriver"))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Retorna sucesso com o response do serviço
	logger.Info("CreateDriver executed successfully",
		zap.String("driverId", driverResponse.Id),
		zap.String("journey", "createDriver"),
	)
	c.JSON(http.StatusCreated, gin.H{
		"message": "Driver created successfully",
		"data":    driverResponse,
	})
}
