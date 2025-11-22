package service

import (
	"github.com/yebology/skillful-certification/app/dto/request"
	"github.com/yebology/skillful-certification/app/dto/response"
	"github.com/yebology/skillful-certification/app/model"
	"github.com/yebology/skillful-certification/app/repository"
)

type ClassParticipantService struct {
	repo *repository.ClassParticipantRepository
}

func NewClassParticipantService(repository *repository.ClassParticipantRepository) *ClassParticipantService {
	return &ClassParticipantService{repo: repository}
}

func (s *ClassParticipantService) AssignParticipantService(dto request.AddClassParticipantDto) error {

	classParticipant := model.ClassParticipant{
		ParticipantId: uint(dto.ParticipantId),
		ClassId:       uint(dto.ClassId),
	}

	return s.repo.Create(&classParticipant)

}

func (s *ClassParticipantService) FetchParticipantClassService(participantId int) ([]response.ParticipantClassDto, error) {

	participantClasses, err := s.repo.FindByParticipantId(participantId)
	if err != nil {
		return []response.ParticipantClassDto{}, err
	}

	var dto []response.ParticipantClassDto

	for _, c := range participantClasses {

		participantClass := response.ParticipantClassDto{
			Id:       c.ID,
			ClassId:  c.ClassId,
			Name:     c.Class.Name,
			Category: c.Class.Category.Name,
		}

		dto = append(dto, participantClass)

	}

	return dto, nil

}

func (s *ClassParticipantService) FetchClassParticipantService(classId int) ([]response.ClassParticipantDto, error) {

	classParticipants, err := s.repo.FindByClassId(classId)
	if err != nil {
		return []response.ClassParticipantDto{}, err
	}

	var dto []response.ClassParticipantDto

	for _, c := range classParticipants {

		classParticipant := response.ClassParticipantDto{
			Id:            c.ID,
			ParticipantId: c.ParticipantId,
			Name:          c.Participant.Name,
			Email:         c.Participant.Email,
			Gender:        c.Participant.Gender.Name,
		}

		dto = append(dto, classParticipant)

	}

	return dto, nil

}

func (s *ClassParticipantService) DeleteClassParticipantService(classParticipantId int) error {

	return s.repo.Delete(classParticipantId)

}
