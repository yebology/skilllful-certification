package service

import (
	"github.com/stretchr/testify/mock"
	"github.com/yebology/skillful-certification/app/dto/request"
	"github.com/yebology/skillful-certification/app/dto/response"
)

type MockClassService struct {
	*ClassService // embed agar bertipe *ClassService
	mock.Mock
}

func (m *MockClassService) CreateClassService(dto request.ClassDto) error {
	args := m.Called(dto)
	return args.Error(0)
}

func (m *MockClassService) EditClassService(id int, dto request.ClassDto) error {
	args := m.Called(id, dto)
	return args.Error(0)
}

func (m *MockClassService) GetAllClassService() ([]response.ClassDto, error) {
	args := m.Called()
	return args.Get(0).([]response.ClassDto), args.Error(1)
}

func (m *MockClassService) GetClassDetailService(id int) (response.ClassDetailDto, error) {
	args := m.Called(id)
	return args.Get(0).(response.ClassDetailDto), args.Error(1)
}

func (m *MockClassService) DeleteClassService(id int) error {
	args := m.Called(id)
	return args.Error(0)
}
