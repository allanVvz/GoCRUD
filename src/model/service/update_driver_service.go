package service

import (
	"errors"

	"github.com/allanVvz/GoCRUD/src/model/repository"
)

type DriverService struct {
	Repo *repository.DriverRepository
}

func (s *DriverService) UpdateDriverStatus(driverID string) error {
	// Busca o motorista pelo ID
	driver, err := s.Repo.FindDriverByID(driverID)
	if err != nil {
		return errors.New("driver not found")
	}

	// Verifica se o status já é 1
	if driver.Status == 1 {
		return errors.New("driver status is already 1")
	}

	// Atualiza o status para 1
	err = s.Repo.UpdateDriverStatus(driverID, 1)
	if err != nil {
		return errors.New("failed to update driver status")
	}

	return nil
}
