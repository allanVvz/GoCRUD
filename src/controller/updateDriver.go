package controller

import (
	"net/http"

	"github.com/allanVvz/GoCRUD/src/config/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (dc *DriverController) UpdateDriver(c *gin.Context) {
	logger.Info("Init UpdateDriver controller",
		zap.String("journey", "updateDriver"),
	)

	// Obter o ID do motorista a partir dos parâmetros da URL
	driverID := c.Param("driverId")
	if driverID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Driver ID is required"})
		return
	}

	// Chama o serviço para atualizar o status do motorista
	err := dc.Service.UpdateDriverStatus(driverID)

	if err != nil {
		logger.Error("Error updating driver status", err,
			zap.String("journey", "updateDriver"),
			zap.String("driverId", driverID),
		)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Retorna sucesso
	logger.Info("UpdateDriver executed successfully",
		zap.String("driverId", driverID),
		zap.String("journey", "updateDriver"),
	)
	c.JSON(http.StatusOK, gin.H{
		"message": "Driver status updated successfully",
	})
}
