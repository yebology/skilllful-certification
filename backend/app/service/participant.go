package service

import (
	"github.com/yebology/skillful-certification/app/dto/request"
	"github.com/yebology/skillful-certification/app/dto/response"
	"github.com/yebology/skillful-certification/app/model"
	"github.com/yebology/skillful-certification/app/repository"
	"github.com/yebology/skillful-certification/utils"
)

type ParticipantService struct {
	repo *repository.ParticipantRepository
}

func NewParticipantService(repository *repository.ParticipantRepository) *ParticipantService {
	return &ParticipantService{repo: repository}
}

func (s *ParticipantService) AddParticipantService(dto request.ParticipantDto) error {

	date, err := utils.ConvertStrToDate(dto.BirthDate)
	if err != nil {
		return err
	}

	participant := model.Participant{
		Name:        dto.Name,
		Email:       dto.Email,
		BirthDate:   date,
		GenderId:    dto.GenderId,
		PhoneNumber: dto.PhoneNumber,
	}

	return s.repo.Create(&participant)

}

func (s *ParticipantService) EditParticipantService(participantId int, dto request.ParticipantDto) error {

	date, err := utils.ConvertStrToDate(dto.BirthDate)
	if err != nil {
		return err
	}

	participant := model.Participant{
		Name:        dto.Name,
		Email:       dto.Email,
		BirthDate:   date,
		GenderId:    dto.GenderId,
		PhoneNumber: dto.PhoneNumber,
	}

	return s.repo.Update(participantId, &participant)

}

func (s *ParticipantService) GetAllParticipantService() ([]response.ParticipantDto, error) {

	participants, err := s.repo.GetAll()
	if err != nil {
		return []response.ParticipantDto{}, err
	}

	var participantsDto []response.ParticipantDto

	for _, p := range participants {

		participant := response.ParticipantDto{
			Id:     p.ID,
			Name:   p.Name,
			Email:  p.Email,
			Gender: p.Gender.Name,
		}

		participantsDto = append(participantsDto, participant)

	}

	return participantsDto, nil

}

func (s *ParticipantService) GetParticipantDetailService(participantId int) (response.ParticipantDetailDto, error) {

	participant, err := s.repo.GetDetail(participantId)
	if err != nil {
		return response.ParticipantDetailDto{}, err
	}

	participantDto := response.ParticipantDetailDto{
		Id:          participant.ID,
		Name:        participant.Name,
		Email:       participant.Email,
		Gender:      participant.Gender.Name,
		BirthDate:   utils.ConvertDateToStr(participant.BirthDate),
		PhoneNumber: participant.PhoneNumber,
	}

	return participantDto, nil

}

func (s *ParticipantService) DeleteParticipantService(participantId int) error {

	return s.repo.Delete(participantId)

}
