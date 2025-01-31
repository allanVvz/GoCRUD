package mock

import (
	"github.com/allanVvz/GoCRUD/src/controller/model/request"
	"github.com/allanVvz/GoCRUD/src/config/rest_err"
	"github.com/stretchr/testify/mock"
)

// Mock para o repositório de jornadas
type MockJourneyRepository struct {
	mock.Mock
}

func (m *MockJourneyRepository) CreateJourney(journey request.JourneyRequest) (*rest_err.RestErr) {
	args := m.Called(journey)
	return args.Get(0).(*rest_err.RestErr)
}
