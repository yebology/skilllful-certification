package service

import (
	"github.com/stretchr/testify/mock"
	"github.com/yebology/skillful-certification/app/dto/request"
	"github.com/yebology/skillful-certification/app/dto/response"
)

type MockClassParticipantService struct {
	mock.Mock
}

func (m *MockClassParticipantService) AssignParticipantService(dto request.AddClassParticipantDto) error {
	args := m.Called(dto)
	return args.Error(0)
}

func (m *MockClassParticipantService) FetchParticipantClassService(participantId int) ([]response.ParticipantClassDto, error) {
	args := m.Called(participantId)
	return args.Get(0).([]response.ParticipantClassDto), args.Error(1)
}

func (m *MockClassParticipantService) FetchClassParticipantService(classId int) ([]response.ClassParticipantDto, error) {
	args := m.Called(classId)
	return args.Get(0).([]response.ClassParticipantDto), args.Error(1)
}

func (m *MockClassParticipantService) DeleteClassParticipantService(classParticipantId int) error {
	args := m.Called(classParticipantId)
	return args.Error(0)
}
