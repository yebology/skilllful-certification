package seed

import (
	"log"

	"github.com/yebology/skillful-certification/app/model"
	"github.com/yebology/skillful-certification/constant"
	"gorm.io/gorm"
)

func SeedGenders(db *gorm.DB) {

	genders := []model.Gender{
		{Name: "Laki-laki"},
		{Name: "Wanita"},
	}

	err := db.Create(&genders).Error
	if err != nil {
		log.Fatalf(constant.ErrSeedingDatabase)
	}

}
