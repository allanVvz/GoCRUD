package controller

import (
	"net/http"

	"github.com/allanVvz/GoCRUD/src/config/logger"
	"github.com/allanVvz/GoCRUD/src/config/rest_err"
	"github.com/allanVvz/GoCRUD/src/controller/model/request"
	"github.com/allanVvz/GoCRUD/src/model/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DriverController struct {
	Service *service.DriverService
}

func NewDriverController(service *service.DriverService) *DriverController {
	return &DriverController{
		Service: service,
	}
}

func (dc *DriverController) Create_Driver(c *gin.Context) {
	var driverRequest request.DriverRequest

	// Bind JSON no request
	if err := c.ShouldBindJSON(&driverRequest); err != nil {
		logger.Error("Invalid JSON received in CreateDriver", err,
			zap.String("journey", "createDriver"))
		c.JSON(http.StatusBadRequest, rest_err.NewBadRequestError("Invalid JSON format"))
		return
	}

	// Chama o serviço para criar o motorista
	driverResponse, err := dc.Service.CreateDriver(driverRequest)
	if err != nil {
		// Verifica se o erro é do tipo *rest_err.RestErr
		if restErr, ok := err.(*rest_err.RestErr); ok {
			logger.Error("Error calling CreateDriver service", err,
				zap.String("journey", "createDriver"))
			c.JSON(restErr.Code, restErr)
			return
		}

		// Caso seja um erro inesperado, retorna Internal Server Error
		logger.Error("Unexpected error in CreateDriver", err,
			zap.String("journey", "createDriver"))
		c.JSON(http.StatusInternalServerError, rest_err.NewInternalServerError("Unexpected error"))
		return
	}

	// Retorna sucesso com os dados criados
	logger.Info("Driver created successfully", zap.String("journey", "createDriver"))
	c.JSON(http.StatusCreated, driverResponse)
}
