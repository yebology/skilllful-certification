package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/yebology/skillful-certification/app/dto/request"
	"github.com/yebology/skillful-certification/app/dto/response"
	"github.com/yebology/skillful-certification/app/service"
	"github.com/yebology/skillful-certification/constant"
)

func setupClassParticipantApp(mockService service.ClassParticipantServiceInterface) *fiber.App {
	app := fiber.New()
	ctrl := NewClassParticipantController(mockService)

	app.Post("/class-participant", ctrl.AssignParticipant)
	app.Get("/participant/:participant_id/classes", ctrl.GetParticipantClass)
	app.Get("/class/:class_id/participants", ctrl.GetClassParticipant)
	app.Delete("/class-participant/:id", ctrl.DeleteClassParticipant)

	return app
}

func TestAssignParticipant_Success(t *testing.T) {
	dto := request.AddClassParticipantDto{
		ParticipantId: 1,
		ClassId:       1,
	}

	mockService := &service.MockClassParticipantService{}

	mockService.On("AssignParticipantService", dto).Return(nil)

	app := setupClassParticipantApp(mockService)

	body, _ := json.Marshal(dto)
	req := httptest.NewRequest("POST", "/class-participant", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestAssignParticipant_ErrorBodyInvalid(t *testing.T) {
	mockService := &service.MockClassParticipantService{}
	app := setupClassParticipantApp(mockService)

	dto := request.AddClassParticipantDto{ParticipantId: 0, ClassId: 1} // invalid participantId
	body, _ := json.Marshal(dto)

	req := httptest.NewRequest("POST", "/class-participant", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
}

func TestGetParticipantClass_Success(t *testing.T) {
	mockService := &service.MockClassParticipantService{}
	mockService.On("FetchParticipantClassService", 1).Return([]response.ParticipantClassDto{
		{Id: 1, ClassId: 1, Name: "Class1", Category: "Category1"},
	}, nil)

	app := setupClassParticipantApp(mockService)
	req := httptest.NewRequest("GET", "/participant/1/classes", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestGetParticipantClass_NotFoundError(t *testing.T) {
	mockService := &service.MockClassParticipantService{}
	mockService.On("FetchParticipantClassService", 1).Return([]response.ParticipantClassDto{}, fmt.Errorf(constant.ErrFetchParticipantClass))

	app := setupClassParticipantApp(mockService)
	req := httptest.NewRequest("GET", "/participant/1/classes", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestGetClassParticipant_Success(t *testing.T) {
	mockService := &service.MockClassParticipantService{}
	mockService.On("FetchClassParticipantService", 1).Return([]response.ClassParticipantDto{
		{Id: 1, ParticipantId: 1, Name: "John Doe", Email: "john@example.com", Gender: "Male"},
	}, nil)

	app := setupClassParticipantApp(mockService)
	req := httptest.NewRequest("GET", "/class/1/participants", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestGetClassParticipant_NotFoundError(t *testing.T) {
	mockService := &service.MockClassParticipantService{}
	mockService.On("FetchClassParticipantService", 1).Return([]response.ClassParticipantDto{}, fmt.Errorf(constant.ErrFetchClassParticipants))

	app := setupClassParticipantApp(mockService)
	req := httptest.NewRequest("GET", "/class/1/participants", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestDeleteClassParticipant_Success(t *testing.T) {
	mockService := &service.MockClassParticipantService{}
	mockService.On("DeleteClassParticipantService", 1).Return(nil)

	app := setupClassParticipantApp(mockService)
	req := httptest.NewRequest("DELETE", "/class-participant/1", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestDeleteClassParticipant_NotFoundError(t *testing.T) {
	mockService := &service.MockClassParticipantService{}
	mockService.On("DeleteClassParticipantService", 1).Return(fmt.Errorf(constant.ErrDeleteEnrollment))

	app := setupClassParticipantApp(mockService)
	req := httptest.NewRequest("DELETE", "/class-participant/1", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	mockService.AssertExpectations(t)
}
