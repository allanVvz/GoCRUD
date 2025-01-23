package service

import (
	"errors"

	"github.com/allanVvz/GoCRUD/src/controller/model/request"
	"github.com/allanVvz/GoCRUD/src/controller/model/response"
	"github.com/allanVvz/GoCRUD/src/model"
	"github.com/allanVvz/GoCRUD/src/model/repository"
)

// DriverService centraliza a lógica de negócio para o recurso Driver.
type DriverService struct {
	Repo *repository.DriverRepository
}

// NewDriverService cria uma nova instância de DriverService.
func NewDriverService(repo *repository.DriverRepository) *DriverService {
	return &DriverService{
		Repo: repo,
	}
}

// CreateDriver cria um novo motorista.
func (s *DriverService) CreateDriver(req request.DriverRequest) (response.DriverResponse, error) {
	// Validações de negócio
	if req.Salary < 1000 {
		return response.DriverResponse{}, errors.New("salary must be at least 1000")
	}

	// Converte o request para o modelo de domínio
	driver := model.Driver{
		Id:           req.Id,
		Name:         req.Name,
		Rg:           req.Rg,
		Registration: req.Registration,
		Salary:       req.Salary,
	}

	// Retorna a resposta com os dados do motorista criado
	return response.DriverResponse{
		Id:           driver.Id,
		Name:         driver.Name,
		Rg:           driver.Rg,
		Registration: driver.Registration,
		Salary:       driver.Salary,
	}, nil
}
