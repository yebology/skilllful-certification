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
	"github.com/stretchr/testify/mock"
	"github.com/yebology/skillful-certification/app/dto/request"
	"github.com/yebology/skillful-certification/app/dto/response"
	"github.com/yebology/skillful-certification/app/service"
	"github.com/yebology/skillful-certification/constant"
)

func setupClassApp(mockService *service.MockClassService) (*fiber.App, *ClassController) {
	app := fiber.New()
	ctrl := NewClassController(mockService)

	app.Post("/class", ctrl.CreateClass)
	app.Put("/class/:id", ctrl.EditClass)
	app.Delete("/class/:id", ctrl.DeleteClass)
	app.Get("/class/:id", ctrl.GetClassDetail)
	app.Get("/class", ctrl.GetAllClass)

	return app, ctrl
}

func TestCreateClass_Success(t *testing.T) {
	mockService := &service.MockClassService{ClassService: &service.ClassService{}}
	mockService.On("CreateClassService", mock.Anything).Return(nil)

	app, _ := setupClassApp(mockService)

	dto := request.ClassDto{
		Name:        "Test Class",
		Description: "Test Description",
		Instructor:  "John Doe",
		CategoryId:  1,
	}
	body, _ := json.Marshal(dto)

	req := httptest.NewRequest("POST", "/class", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	mockService.AssertExpectations(t)
}

func TestCreateClass_ErrorBodyInvalid(t *testing.T) {
	mockService := &service.MockClassService{}

	app, _ := setupClassApp(mockService)

	dto := request.ClassDto{
		Name:        "",
		Description: "Test Description",
		Instructor:  "John Doe",
		CategoryId:  1,
	}
	body, _ := json.Marshal(dto)

	req := httptest.NewRequest("POST", "/class", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)
	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
}

func TestEditClass_Success(t *testing.T) {
	mockService := &service.MockClassService{}
	mockService.On("EditClassService", mock.Anything, mock.Anything).Return(nil)

	app, _ := setupClassApp(mockService)

	dto := request.ClassDto{
		Name:        "Edited Class",
		Description: "Updated Description",
		Instructor:  "Jane Doe",
		CategoryId:  2,
	}
	body, _ := json.Marshal(dto)

	req := httptest.NewRequest("PUT", "/class/1", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	mockService.AssertExpectations(t)
}

func TestDeleteClass_Success(t *testing.T) {
	mockService := &service.MockClassService{ClassService: &service.ClassService{}}
	mockService.On("DeleteClassService", 1).Return(nil)

	app, _ := setupClassApp(mockService)

	req := httptest.NewRequest("DELETE", "/class/1", nil)
	resp, _ := app.Test(req)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	mockService.AssertExpectations(t)
}

func TestDeleteClass_NotFoundError(t *testing.T) {
	mockService := &service.MockClassService{}
	mockService.On("DeleteClassService", 1).Return(fmt.Errorf(constant.ErrDeleteClass))

	app, _ := setupClassApp(mockService)
	req := httptest.NewRequest("DELETE", "/class/1", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestGetClassDetail_Success(t *testing.T) {
	mockService := &service.MockClassService{}
	mockService.On("GetClassDetailService", mock.Anything).Return(response.ClassDetailDto{
		Id:          1,
		Name:        "Test Class",
		Description: "Test Description",
		Instructor:  "John Doe",
		Category:    "Category 1",
	}, nil)

	app, _ := setupClassApp(mockService)

	req := httptest.NewRequest("GET", "/class/1", nil)
	resp, _ := app.Test(req)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	mockService.AssertExpectations(t)
}

func TestGetClassDetail_NotFound(t *testing.T) {
	mockService := &service.MockClassService{}

	mockService.On("GetClassDetailService", 1).Return(response.ClassDetailDto{}, fmt.Errorf(constant.ErrFetchClass))

	app, _ := setupClassApp(mockService)

	req := httptest.NewRequest("GET", "/class/1", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestGetAllClass_Success(t *testing.T) {
	mockService := &service.MockClassService{}

	mockService.On("GetAllClassService").Return([]response.ClassDto{
		{Id: 1, Name: "Class1", Category: "Category1"},
		{Id: 2, Name: "Class2", Category: "Category2"},
	}, nil)

	app, _ := setupClassApp(mockService)

	req := httptest.NewRequest("GET", "/class", nil)
	resp, _ := app.Test(req)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	mockService.AssertExpectations(t)
}
