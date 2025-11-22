package seed

import (
	"log"

	"github.com/yebology/skillful-certification/app/model"
	"github.com/yebology/skillful-certification/constant"
	"gorm.io/gorm"
)

func SeedEnrollments(db *gorm.DB) {

	var enrollments []model.ClassParticipant

	for i := 0; i < 4; i++ {

		enrollment := model.ClassParticipant{
			ClassId:       uint(i + 1),
			ParticipantId: uint(i + 1),
		}

		enrollments = append(enrollments, enrollment)
	}

	err := db.Create(&enrollments).Error
	if err != nil {
		log.Fatalf(constant.ErrSeedingDatabase)
	}

}
