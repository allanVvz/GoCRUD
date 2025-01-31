package mock

import (
	"github.com/allanVvz/GoCRUD/src/controller/model/request"
	"github.com/allanVvz/GoCRUD/src/controller/model/response"
	"github.com/stretchr/testify/mock"
)

// Mock para Jouneys
type MockJourneyService struct {
	mock.Mock
}

func (m *MockJourneyService) CreateJourney(journeyRequest request.JourneyRequest) (response.JourneyResponse, error) {
	args := m.Called(journeyRequest)
	return args.Get(0).(response.JourneyResponse), args.Error(1)
}
