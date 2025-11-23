package service

import (
	"github.com/stretchr/testify/mock"
	"github.com/yebology/skillful-certification/app/dto/request"
	"github.com/yebology/skillful-certification/app/dto/response"
)

type MockParticipantService struct {
	*ParticipantService
	mock.Mock
}

func (m *MockParticipantService) AddParticipantService(dto request.ParticipantDto) error {
	args := m.Called(dto)
	return args.Error(0)
}

func (m *MockParticipantService) EditParticipantService(id int, dto request.ParticipantDto) error {
	args := m.Called(id, dto)
	return args.Error(0)
}

func (m *MockParticipantService) GetAllParticipantService() ([]response.ParticipantDto, error) {
	args := m.Called()
	return args.Get(0).([]response.ParticipantDto), args.Error(1)
}

func (m *MockParticipantService) GetParticipantDetailService(id int) (response.ParticipantDetailDto, error) {
	args := m.Called(id)
	return args.Get(0).(response.ParticipantDetailDto), args.Error(1)
}

func (m *MockParticipantService) DeleteParticipantService(id int) error {
	args := m.Called(id)
	return args.Error(0)
}