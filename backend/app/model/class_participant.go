package model

import "gorm.io/gorm"

type ClassParticipant struct {
	gorm.Model
	ClassId       uint
	Class         Class `gorm:"foreignKey:ClassId;constraint:OnDelete:CASCADE;"`
	ParticipantId uint
	Participant   Participant `gorm:"foreignKey:ParticipantId;constraint:OnDelete:CASCADE;"`
}
