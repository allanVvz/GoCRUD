package service

import (
	"github.com/allanVvz/GoCRUD/src/config/logger"
	"github.com/allanVvz/GoCRUD/src/config/rest_err"
	"go.uber.org/zap"
)

func (s *DriverService) FireDriver(driverID string) *rest_err.RestErr {
	logger.Info("Starting UpdateDriverStatus service", zap.String("driverID", driverID))

	// Busca o motorista pelo ID
	driver, err := s.Repo.FindDriverByID(driverID)
	if err != nil {
		logger.Error("Error finding driver in UpdateDriverStatus", err, zap.String("driverID", driverID))
		if err.Error() == "driver not found" {
			return rest_err.NewNotFoundError("Driver not found")
		}
		return rest_err.NewInternalServerError("Error finding driver")
	}

	// Verifica se o status já é 0
	if !driver.Status {
		logger.Info("Driver status is already 0, no update needed", zap.String("driverID", driverID))
		return rest_err.NewBadRequestError("Driver status is already inactive (0)")
	}

	// Atualiza o status para 0
	err = s.Repo.FireDriver(driverID, false)
	if err != nil {
		logger.Error("Error updating driver status in UpdateDriverStatus", err, zap.String("driverID", driverID))
		return rest_err.NewInternalServerError("Error updating driver status")
	}

	logger.Info("Driver status updated successfully", zap.String("driverID", driverID))
	return nil
}
