package service

import (
	"github.com/allanVvz/GoCRUD/src/config/logger"
	"github.com/allanVvz/GoCRUD/src/config/rest_err"
	"github.com/allanVvz/GoCRUD/src/controller/model/response"
	"go.uber.org/zap"
)

// FindDriverByID busca um motorista pelo ID e retorna um response.DriverResponse.
func (s *DriverService) FindDriverByID(driverID string) (response.DriverResponse, *rest_err.RestErr) {
	logger.Info("Starting FindDriverByID service", zap.String("driverID", driverID))

	// Chama o repository para buscar o motorista pelo ID
	driver, err := s.Repo.FindDriverByID(driverID)
	if err != nil {
		logger.Error("Error finding driver in FindDriverByID", err, zap.String("driverID", driverID))
		return response.DriverResponse{}, err
	}

	// Converte `request.DriverRequest` para `response.DriverResponse`
	return response.DriverResponse(driver), nil
}

// FindDriverByReg busca um motorista pelo registro e retorna um response.DriverResponse.
func (s *DriverService) FindDriverByReg(driverReg string) (response.DriverResponse, *rest_err.RestErr) {
	logger.Info("Starting FindDriverByReg service", zap.String("driverReg", driverReg))

	// Chama o repository para buscar o motorista pelo registro
	driver, err := s.Repo.FindDriverByReg(driverReg)
	if err != nil {
		logger.Error("Error finding driver in FindDriverByReg", err, zap.String("driverReg", driverReg))
		return response.DriverResponse{}, err
	}

	// Converte `request.DriverRequest` para `response.DriverResponse`
	return response.DriverResponse(driver), nil
}

// FindDriverByRg busca um motorista pelo RG e retorna um response.DriverResponse.
func (s *DriverService) FindDriverByRg(driverRg string) (response.DriverResponse, *rest_err.RestErr) {
	logger.Info("Starting FindDriverByRg service", zap.String("driverRg", driverRg))

	// Chama o repository para buscar o motorista pelo RG
	driver, err := s.Repo.FindDriverByRg(driverRg)
	if err != nil {
		logger.Error("Error finding driver in FindDriverByRg", err, zap.String("driverRg", driverRg))
		return response.DriverResponse{}, err
	}

	// Converte `request.DriverRequest` para `response.DriverResponse`
	return response.DriverResponse(driver), nil
}
