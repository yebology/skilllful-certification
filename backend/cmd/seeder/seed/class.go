package seed

import (
	"log"
	"math/rand/v2"

	"github.com/bxcodec/faker/v3"
	"github.com/yebology/skillful-certification/app/model"
	"github.com/yebology/skillful-certification/constant"
	"gorm.io/gorm"
)

func SeedClasses(db *gorm.DB) {

	var classes []model.Class

	for i := 0; i < 4; i++ {

		class := model.Class{
			Name:        faker.Word(),
			Description: faker.Paragraph(),
			Instructor:  faker.Name(),
			CategoryId:  uint(rand.IntN(4) + 1),
		}

		classes = append(classes, class)
	}

	err := db.Create(&classes).Error
	if err != nil {
		log.Fatalf(constant.ErrSeedingDatabase)
	}

}
