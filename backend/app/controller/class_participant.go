package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yebology/skillful-certification/app/dto/request"
	"github.com/yebology/skillful-certification/app/service"
	"github.com/yebology/skillful-certification/constant"
	"github.com/yebology/skillful-certification/output"
	"github.com/yebology/skillful-certification/utils"
)

type ClassParticipantController struct {
	service service.ClassParticipantServiceInterface
}

func NewClassParticipantController(service service.ClassParticipantServiceInterface) *ClassParticipantController {
	return &ClassParticipantController{service: service}
}

func (con *ClassParticipantController) AssignParticipant(c *fiber.Ctx) error {

	var dto request.AddClassParticipantDto

	err := utils.ParseAndValidateBody(c, &dto)
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusInternalServerError, constant.ErrParsingBody, nil)
	}

	err = con.service.AssignParticipantService(dto)
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusInternalServerError, err.Error(), nil)
	}

	return output.GetOutput(c, constant.StatusSuccess, fiber.StatusOK, constant.SuccessAssignParticipantToClass, nil)

}

func (con *ClassParticipantController) GetParticipantClass(c *fiber.Ctx) error {

	participantId, err := utils.ConvertToNum(c, "participant_id")
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusInternalServerError, constant.ErrConvertNum, nil)
	}

	classes, err := con.service.FetchParticipantClassService(participantId)
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusNotFound, err.Error(), nil)
	}

	return output.GetOutput(c, constant.StatusSuccess, fiber.StatusOK, constant.SuccessFetchParticipantClass, classes)

}

func (con *ClassParticipantController) GetClassParticipant(c *fiber.Ctx) error {

	classId, err := utils.ConvertToNum(c, "class_id")
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusInternalServerError, constant.ErrConvertNum, nil)
	}

	participants, err := con.service.FetchClassParticipantService(classId)
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusNotFound, err.Error(), nil)
	}

	return output.GetOutput(c, constant.StatusSuccess, fiber.StatusOK, constant.SuccessFetchClassParticipants, participants)

}

func (con *ClassParticipantController) DeleteClassParticipant(c *fiber.Ctx) error {

	classParticipantId, err := utils.ConvertToNum(c, "id")
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusInternalServerError, constant.ErrConvertNum, nil)
	}

	err = con.service.DeleteClassParticipantService(classParticipantId)
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusNotFound, err.Error(), nil)
	}

	return output.GetOutput(c, constant.StatusSuccess, fiber.StatusOK, constant.SuccessDeleteEnrollment, nil)

}
