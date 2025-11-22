package seed

import (
	"log"

	"github.com/yebology/skillful-certification/app/model"
	"github.com/yebology/skillful-certification/constant"
	"gorm.io/gorm"
)

func SeedCategories(db *gorm.DB) {

	categories := []model.Category{
		{Name: "Desain Grafis"},
		{Name: "Pemrograman Dasar"},
		{Name: "Editing Video"},
		{Name: "Public Speaking"},
	}

	err := db.Create(&categories).Error
	if err != nil {
		log.Fatalf(constant.ErrSeedingDatabase)
	}

}
