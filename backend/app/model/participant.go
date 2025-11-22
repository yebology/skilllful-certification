package model

import (
	"time"

	"gorm.io/gorm"
)

type Participant struct {
	gorm.Model
	Name             string
	Email            string
	PhoneNumber      string
	BirthDate        time.Time
	GenderId         uint
	Gender           Gender             `gorm:"foreignKey:GenderId;constraint:OnDelete:SET NULL;"`
	ClassParticipant []ClassParticipant `gorm:"foreignKey:ParticipantId;constraint:OnDelete:CASCADE;"`
}
