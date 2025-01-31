package test

import (
	"testing"

	"github.com/allanVvz/GoCRUD/src/config/rest_err"
	"github.com/allanVvz/GoCRUD/src/controller/model/request"
	"github.com/allanVvz/GoCRUD/src/controller/model/response"
	"github.com/allanVvz/GoCRUD/src/model/repository/mock"
	"github.com/stretchr/testify/assert"
)

// Testa a criação de uma jornada no serviço
func TestCreateJourney(t *testing.T) {
	mockRepo := new(mock.MockJourneyRepository)

	newJourney := response.JourneyResponse{
		Id:              "987654",
		DriverId:        "driver567",
		VehicleId:       "vehicle890",
		InitialDate:     "2025-02-01",
		FinalDate:       "2025-02-05",
		InitialLocation: "City A",
		FinalLocation:   "City B",
		Duration:        15.2,
	}

	newJourneyRequest := request.JourneyRequest{
		Id:              newJourney.Id,
		DriverID:        newJourney.DriverId,
		VehicleID:       newJourney.VehicleId,
		InitialDate:     newJourney.InitialDate,
		FinalDate:       newJourney.FinalDate,
		InitialLocation: newJourney.InitialLocation,
		FinalLocation:   newJourney.FinalLocation,
	}

	// Configura o mock para retornar sucesso ao criar uma jornada
	mockRepo.On("CreateJourney", newJourneyRequest).Return((*rest_err.RestErr)(nil))

	// Executa o teste chamando o serviço
	err := mockRepo.CreateJourney(newJourneyRequest)

	// Verifica os resultados
	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}
