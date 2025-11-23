package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yebology/skillful-certification/app/dto/request"
	"github.com/yebology/skillful-certification/app/service"
	"github.com/yebology/skillful-certification/constant"
	"github.com/yebology/skillful-certification/output"
	"github.com/yebology/skillful-certification/utils"
)

type ClassController struct {
	service service.ClassServiceInterface
}

func NewClassController(s service.ClassServiceInterface) *ClassController {
	return &ClassController{service: s}
}

func (con *ClassController) CreateClass(c *fiber.Ctx) error {

	var dto request.ClassDto

	err := utils.ParseAndValidateBody(c, &dto)
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusInternalServerError, constant.ErrParsingBody, nil)
	}

	err = con.service.CreateClassService(dto)
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusInternalServerError, err.Error(), nil)
	}

	return output.GetOutput(c, constant.StatusSuccess, fiber.StatusOK, constant.SuccessCreateClass, nil)

}

func (con *ClassController) EditClass(c *fiber.Ctx) error {

	var dto request.ClassDto

	err := utils.ParseAndValidateBody(c, &dto)
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusInternalServerError, constant.ErrParsingBody, nil)
	}

	classId, err := utils.ConvertToNum(c, "id")
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusInternalServerError, constant.ErrConvertNum, nil)
	}

	err = con.service.EditClassService(classId, dto)
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusInternalServerError, err.Error(), nil)
	}

	return output.GetOutput(c, constant.StatusSuccess, fiber.StatusOK, constant.SuccessEditClass, nil)

}

func (con *ClassController) GetAllClass(c *fiber.Ctx) error {

	classes, err := con.service.GetAllClassService()
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusInternalServerError, err.Error(), nil)
	}

	return output.GetOutput(c, constant.StatusSuccess, fiber.StatusOK, constant.SuccessFetchClass, classes)

}

func (con *ClassController) GetClassDetail(c *fiber.Ctx) error {

	classId, err := utils.ConvertToNum(c, "id")
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusInternalServerError, constant.ErrConvertNum, nil)
	}

	class, err := con.service.GetClassDetailService(classId)
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusNotFound, err.Error(), nil)
	}

	return output.GetOutput(c, constant.StatusSuccess, fiber.StatusOK, constant.SuccessFetchClass, class)

}

func (con *ClassController) DeleteClass(c *fiber.Ctx) error {

	classId, err := utils.ConvertToNum(c, "id")
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusInternalServerError, constant.ErrConvertNum, nil)
	}

	err = con.service.DeleteClassService(classId)
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusNotFound, err.Error(), nil)
	}

	return output.GetOutput(c, constant.StatusSuccess, fiber.StatusOK, constant.SuccessDeleteClass, nil)

}
