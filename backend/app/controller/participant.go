package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yebology/skillful-certification/app/dto/request"
	"github.com/yebology/skillful-certification/app/service"
	"github.com/yebology/skillful-certification/constant"
	"github.com/yebology/skillful-certification/output"
	"github.com/yebology/skillful-certification/utils"
)

type ParticipantController struct {
	service service.ParticipantServiceInterface
}

func NewParticipantController(s service.ParticipantServiceInterface) *ParticipantController {
	return &ParticipantController{service: s}
}

func (con *ParticipantController) AddParticipant(c *fiber.Ctx) error {

	var dto request.ParticipantDto

	err := utils.ParseAndValidateBody(c, &dto)
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusInternalServerError, constant.ErrParsingBody, nil)
	}

	err = con.service.AddParticipantService(dto)
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusInternalServerError, err.Error(), nil)
	}

	return output.GetOutput(c, constant.StatusSuccess, fiber.StatusOK, constant.SuccessAddParticipant, nil)

}

func (con *ParticipantController) EditParticipant(c *fiber.Ctx) error {

	var dto request.ParticipantDto

	err := utils.ParseAndValidateBody(c, &dto)
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusInternalServerError, constant.ErrParsingBody, nil)
	}

	participantId, err := utils.ConvertToNum(c, "id")
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusInternalServerError, constant.ErrConvertNum, nil)
	}

	err = con.service.EditParticipantService(participantId, dto)
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusInternalServerError, err.Error(), nil)
	}

	return output.GetOutput(c, constant.StatusSuccess, fiber.StatusOK, constant.SuccessEditParticipant, nil)

}

func (con *ParticipantController) GetAllParticipant(c *fiber.Ctx) error {

	participants, err := con.service.GetAllParticipantService()
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusInternalServerError, err.Error(), nil)
	}

	return output.GetOutput(c, constant.StatusSuccess, fiber.StatusOK, constant.SuccessFetchParticipant, participants)

}

func (con *ParticipantController) GetParticipantDetail(c *fiber.Ctx) error {

	participantId, err := utils.ConvertToNum(c, "id")
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusInternalServerError, constant.ErrConvertNum, nil)
	}

	participant, err := con.service.GetParticipantDetailService(participantId)
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusInternalServerError, err.Error(), nil)
	}

	return output.GetOutput(c, constant.StatusSuccess, fiber.StatusOK, constant.SuccessFetchParticipant, participant)

}

func (con *ParticipantController) DeleteParticipant(c *fiber.Ctx) error {

	participantId, err := utils.ConvertToNum(c, "id")
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusInternalServerError, constant.ErrConvertNum, nil)
	}

	err = con.service.DeleteParticipantService(participantId)
	if err != nil {
		return output.GetOutput(c, constant.StatusError, fiber.StatusNotFound, err.Error(), nil)
	}

	return output.GetOutput(c, constant.StatusSuccess, fiber.StatusOK, constant.SuccessDeleteParticipant, nil)

}
