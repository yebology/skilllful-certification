package repository

import (
	"fmt"

	"github.com/yebology/skillful-certification/app/model"
	"github.com/yebology/skillful-certification/constant"
	"gorm.io/gorm"
)

type ParticipantRepository struct {
	db *gorm.DB
}

func NewParticipantRepository(db *gorm.DB) *ParticipantRepository {
	return &ParticipantRepository{db: db}
}

func (r *ParticipantRepository) Create(participant *model.Participant) error {

	err := r.db.Create(participant).Error
	if err != nil {
		return fmt.Errorf(constant.ErrAddParticipant)
	}

	return nil
}

func (r *ParticipantRepository) GetAll() ([]model.Participant, error) {

	var participants []model.Participant

	err := r.db.Preload("Gender").
		Find(&participants).
		Error
	if err != nil {
		return []model.Participant{}, fmt.Errorf(constant.ErrFetchParticipant)
	}

	return participants, nil

}

func (r *ParticipantRepository) GetDetail(participantId int) (model.Participant, error) {

	var participant model.Participant

	err := r.db.Preload("Gender").
		First(&participant, participantId).
		Error
	if err != nil {
		return model.Participant{}, fmt.Errorf(constant.ErrFetchParticipant)
	}

	return participant, nil

}

func (r *ParticipantRepository) Update(participantId int, participant *model.Participant) error {

	err := r.db.Model(&model.Participant{}).
		Where(`id = ?`, participantId).
		Updates(participant).
		Error
	if err != nil {
		return fmt.Errorf(constant.ErrEditParticipant)
	}

	return nil

}

func (r *ParticipantRepository) Delete(participantId int) error {

	err := r.db.Unscoped().Delete(&model.Participant{}, participantId).Error
	if err != nil {
		return fmt.Errorf(constant.ErrDeleteParticipant)
	}

	return nil
}
