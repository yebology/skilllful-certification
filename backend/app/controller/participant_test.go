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

func setupParticipantApp(mockService *service.MockParticipantService) *fiber.App {
	app := fiber.New()
	ctrl := NewParticipantController(mockService)

	app.Post("/participant", ctrl.AddParticipant)
	app.Put("/participant/:id", ctrl.EditParticipant)
	app.Delete("/participant/:id", ctrl.DeleteParticipant)
	app.Get("/participant/:id", ctrl.GetParticipantDetail)
	app.Get("/participant", ctrl.GetAllParticipant)
	return app
}

func TestAddParticipant_Success(t *testing.T) {
	dto := request.ParticipantDto{
		Name:        "John Doe",
		Email:       "john@example.com",
		BirthDate:   "2000-01-01",
		GenderId:    1,
		PhoneNumber: "08123456789",
	}

	mockService := &service.MockParticipantService{}

	mockService.On("AddParticipantService", dto).Return(nil)

	app := setupParticipantApp(mockService)

	body, _ := json.Marshal(dto)
	req := httptest.NewRequest("POST", "/participant", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	mockService.AssertExpectations(t)
}

func TestAddParticipant_InvalidBody(t *testing.T) {
	mockService := &service.MockParticipantService{}
	app := setupParticipantApp(mockService)

	dto := request.ParticipantDto{Name: ""}
	body, _ := json.Marshal(dto)

	req := httptest.NewRequest("POST", "/participant", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)
	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
}

func TestGetParticipantDetail_Success(t *testing.T) {
	mockService := &service.MockParticipantService{}
	mockService.On("GetParticipantDetailService", 1).Return(response.ParticipantDetailDto{
		Id:          1,
		Name:        "John Doe",
		Email:       "john@example.com",
		Gender:      "Male",
		BirthDate:   "2000-01-01",
		PhoneNumber: "08123456789",
	}, nil)

	app := setupParticipantApp(mockService)

	req := httptest.NewRequest("GET", "/participant/1", nil)
	resp, _ := app.Test(req)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestGetParticipantDetail_NotFound(t *testing.T) {
	mockService := &service.MockParticipantService{}
	mockService.On("GetParticipantDetailService", 1).Return(response.ParticipantDetailDto{}, fmt.Errorf("participant not found"))

	app := setupParticipantApp(mockService)

	req := httptest.NewRequest("GET", "/participant/1", nil)
	resp, _ := app.Test(req)
	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestGetAllParticipant_Success(t *testing.T) {
	mockService := &service.MockParticipantService{}
	mockService.On("GetAllParticipantService").Return([]response.ParticipantDto{
		{Id: 1, Name: "John Doe", Email: "john@example.com", Gender: "Male"},
		{Id: 2, Name: "Jane Doe", Email: "jane@example.com", Gender: "Female"},
	}, nil)

	app := setupParticipantApp(mockService)

	req := httptest.NewRequest("GET", "/participant", nil)
	resp, _ := app.Test(req)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestDeleteParticipant_Success(t *testing.T) {
	mockService := &service.MockParticipantService{}
	mockService.On("DeleteParticipantService", 1).Return(nil)

	app := setupParticipantApp(mockService)

	req := httptest.NewRequest("DELETE", "/participant/1", nil)
	resp, _ := app.Test(req)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestDeleteParticipant_NotFoundError(t *testing.T) {
	mockService := &service.MockParticipantService{}
	mockService.On("DeleteParticipantService", 1).Return(fmt.Errorf(constant.ErrDeleteParticipant))

	app := setupParticipantApp(mockService)

	req := httptest.NewRequest("DELETE", "/participant/1", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestEditParticipant_Success(t *testing.T) {
	participantID := 1
	dto := request.ParticipantDto{
		Name:        "John Edited",
		Email:       "john@example.com",
		BirthDate:   "2000-01-01",
		GenderId:    1,
		PhoneNumber: "08123456789",
	}

	mockService := &service.MockParticipantService{}

	mockService.On("EditParticipantService", participantID, dto).Return(nil)

	app := setupParticipantApp(mockService)

	body, _ := json.Marshal(dto)
	req := httptest.NewRequest("PUT", fmt.Sprintf("/participant/%d", participantID), bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	mockService.AssertExpectations(t)
}
