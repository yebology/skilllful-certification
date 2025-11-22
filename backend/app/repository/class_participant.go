package repository

import (
	"fmt"

	"github.com/yebology/skillful-certification/app/model"
	"github.com/yebology/skillful-certification/constant"
	"gorm.io/gorm"
)

type ClassParticipantRepository struct {
	db *gorm.DB
}

func NewClassParticipantRepository(db *gorm.DB) *ClassParticipantRepository {
	return &ClassParticipantRepository{db: db}
}

func (r *ClassParticipantRepository) Create(classParticipants *model.ClassParticipant) error {

	err := r.db.Create(classParticipants).Error
	if err != nil {
		return fmt.Errorf(constant.ErrAssignParticipantToClass)
	}

	return nil

}

func (r *ClassParticipantRepository) FindByParticipantId(participantId int) ([]model.ClassParticipant, error) {

	var classParticipants []model.ClassParticipant

	err := r.db.Preload("Class.Category").
		Find(&classParticipants).
		Where(`participant_id = ?`, participantId).
		Error
	if err != nil {
		return []model.ClassParticipant{}, fmt.Errorf(constant.ErrFetchParticipantClass)
	}

	return classParticipants, nil

}

func (r *ClassParticipantRepository) FindByClassId(classId int) ([]model.ClassParticipant, error) {

	var classParticipants []model.ClassParticipant

	err := r.db.Preload("Participant.Gender").
		Find(&classParticipants).
		Where(`class_id = ?`, classId).
		Error
	if err != nil {
		return []model.ClassParticipant{}, fmt.Errorf(constant.ErrFetchClassParticipants)
	}

	return classParticipants, nil

}

func (r *ClassParticipantRepository) Delete(classParticipantId int) error {

	err := r.db.Unscoped().Delete(&model.ClassParticipant{}, classParticipantId).Error
	if err != nil {
		return fmt.Errorf(constant.ErrDeleteEnrollment)
	}

	return nil

}
