package clear

import (
	"log"

	"github.com/yebology/skillful-certification/app/model"
	"github.com/yebology/skillful-certification/constant"
	"gorm.io/gorm"
)

func ClearDB(db *gorm.DB) {

	tables := []interface{}{
		&model.ClassParticipant{},
		&model.Participant{},
		&model.Class{},
		&model.Gender{},
		&model.Category{},
	}

	for _, table := range tables {

		err := db.Migrator().DropTable(table)
		if err != nil {
			log.Fatalf(string(constant.ErrResetTable))
		}

	}

}
