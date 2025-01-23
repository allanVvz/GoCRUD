package service

import (
	"github.com/allanVvz/GoCRUD/src/controller/model/request"
	"github.com/allanVvz/GoCRUD/src/controller/model/response"
	"github.com/allanVvz/GoCRUD/src/model/repository"
)

type DriverService struct {
	Repo *repository.DriverRepository
}

func NewDriverService(repo *repository.DriverRepository) *DriverService {
	return &DriverService{
		Repo: repo,
	}
}

func (s *DriverService) CreateDriver(driverRequest request.DriverRequest) (response.DriverResponse, error) {
	id, err := s.Repo.SaveDriver(driverRequest)
	if err != nil {
		return response.DriverResponse{}, err
	}

	return response.DriverResponse{
		Id:           id,
		Status:       driverRequest.Status,
		Name:         driverRequest.Name,
		Rg:           driverRequest.Rg,
		Registration: driverRequest.Registration,
		Salary:       driverRequest.Salary,
	}, nil
}
