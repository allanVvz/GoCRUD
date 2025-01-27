package service

import (
	"github.com/allanVvz/GoCRUD/src/config/logger"
	"github.com/allanVvz/GoCRUD/src/config/rest_err"
	"github.com/allanVvz/GoCRUD/src/controller/model/response"
	"go.uber.org/zap"
)

func (s *DriverService) FindDriverByID(driverID string) (response.DriverResponse, *rest_err.RestErr) {
	logger.Info("Starting FindDriver service", zap.String("driverID", driverID))

	// Busca o motorista pelo ID no repositório
	driver, err := s.Repo.FindDriverByID(driverID)
	if err != nil {
		logger.Error("Error finding driver in FindDriverByID", err, zap.String("driverID", driverID))
		if err.Error() == "driver not found" {
			return response.DriverResponse{}, rest_err.NewNotFoundError("Driver not found")
		}
		return response.DriverResponse{}, rest_err.NewInternalServerError("Error finding driver")
	}

	// Retorna os dados do motorista no formato JSON (Response)
	logger.Info("Driver found successfully", zap.String("driverID", driverID))
	return response.DriverResponse{
		Id:           driver.Id,
		Status:       driver.Status,
		Name:         driver.Name,
		Rg:           driver.Rg,
		Registration: driver.Registration,
		Salary:       driver.Salary,
	}, nil
}
