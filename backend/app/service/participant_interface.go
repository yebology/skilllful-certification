package service

import (
	"github.com/yebology/skillful-certification/app/dto/request"
	"github.com/yebology/skillful-certification/app/dto/response"
)

type ParticipantServiceInterface interface {
	AddParticipantService(dto request.ParticipantDto) error
	EditParticipantService(id int, dto request.ParticipantDto) error
	GetAllParticipantService() ([]response.ParticipantDto, error)
	GetParticipantDetailService(id int) (response.ParticipantDetailDto, error)
	DeleteParticipantService(id int) error
}