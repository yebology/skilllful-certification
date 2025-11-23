package service

import (
	"github.com/yebology/skillful-certification/app/dto/request"
	"github.com/yebology/skillful-certification/app/dto/response"
)

// Interface untuk ClassParticipantService
type ClassParticipantServiceInterface interface {
	AssignParticipantService(dto request.AddClassParticipantDto) error
	FetchParticipantClassService(participantId int) ([]response.ParticipantClassDto, error)
	FetchClassParticipantService(classId int) ([]response.ClassParticipantDto, error)
	DeleteClassParticipantService(classParticipantId int) error
}
