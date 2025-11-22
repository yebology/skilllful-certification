package seed

import (
	"log"
	"math/rand/v2"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/yebology/skillful-certification/app/model"
	"github.com/yebology/skillful-certification/constant"
	"gorm.io/gorm"
)

func SeedParticipants(db *gorm.DB) {

	var participants []model.Participant

	for i := 0; i < 4; i++ {

		seedDate := faker.Date()
		date, err := time.Parse("2006-01-02", seedDate)
		if err != nil {
			log.Fatalf(constant.ErrSeedingDatabase)
		}

		participant := model.Participant{
			Name:        faker.Name(),
			Email:       faker.Email(),
			PhoneNumber: faker.Phonenumber(),
			BirthDate:   date,
			GenderId:    uint(rand.IntN(2) + 1),
		}

		participants = append(participants, participant)
	}

	err := db.Create(&participants).Error
	if err != nil {
		log.Fatalf(constant.ErrSeedingDatabase)
	}

}
