package model

import "gorm.io/gorm"

type Class struct {
	gorm.Model
	Name             string
	Description      string
	Instructor       string
	CategoryId       uint
	Category         Category           `gorm:"foreignKey:CategoryId;constraint:OnDelete:SET NULL;"`
	ClassParticipant []ClassParticipant `gorm:"foreignKey:ClassId;constraint:OnDelete:CASCADE;"`
}
