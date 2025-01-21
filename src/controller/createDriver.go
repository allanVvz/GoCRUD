package controller

import (
	"net/http"

	"github.com/allanVvz/GoCRUD/src/config/logger"
	"github.com/allanVvz/GoCRUD/src/config/validation"
	"github.com/allanVvz/GoCRUD/src/controller/model/request"
	"github.com/allanVvz/GoCRUD/src/controller/model/response"
	"github.com/allanVvz/GoCRUD/src/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (dc *driverControllerInterface) CreateDriver(c *gin.Context) {
	logger.Info("Init CreateDriver controller",
		zap.String("journey", "createDriver"),
	)

	var driverRequest request.DriverRequest

	// Validação do JSON recebido
	if err := c.ShouldBindJSON(&driverRequest); err != nil {
		logger.Error("Error trying to validate driver info", err,
			zap.String("journey", "createDriver"))
		errRest := validation.ValidateDriverError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	// Criação do domínio do motorista
	domain := model.NewDriverDomain(
		driverRequest.Id,
		driverRequest.Name,
		driverRequest.Rg,
		driverRequest.Registration,
		driverRequest.Salary,
	)

	// Chamando o serviço para criar o motorista
	domainResult, err := dc.service.CreateDriverServices(domain)
	if err != nil {
		logger.Error("Error trying to call CreateDriver service", err,
			zap.String("journey", "createDriver"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("CreateDriver controller executed successfully",
		zap.String("driverId", domainResult.GetID()),
		zap.String("journey", "createDriver"))

	// Retornando o objeto criado ao cliente
	c.JSON(http.StatusCreated, response.DriverResponse{
		Id:           domainResult.GetID(),
		Name:         domainResult.GetName(),
		Rg:           domainResult.GetRg(),
		Registration: domainResult.GetRegistration(),
		Salary:       domainResult.GetSalary(),
	})
}
