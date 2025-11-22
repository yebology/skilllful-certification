package main

import (
	"fmt"
	"log"

	"github.com/yebology/skillful-certification/app/model"
	"github.com/yebology/skillful-certification/cmd/seeder/seed"
	"github.com/yebology/skillful-certification/config"
	"github.com/yebology/skillful-certification/constant"
	"gorm.io/gorm"
)

func main() {

	config.LoadEnv()

	config.ConnectDB()
	database := config.GetDB()

	fmt.Println("🚀 Starting database migration and seeding...")

	autoMigrateAndSeed(database, &model.Category{}, "Category", func() { seed.SeedCategories(database) })
	autoMigrateAndSeed(database, &model.Gender{}, "Gender", func() { seed.SeedGenders(database) })
	autoMigrateAndSeed(database, &model.Participant{}, "Participant", func() { seed.SeedParticipants(database) })
	autoMigrateAndSeed(database, &model.Class{}, "Class", func() { seed.SeedClasses(database) })
	autoMigrateAndSeed(database, &model.ClassParticipant{}, "Class Participant", func() { seed.SeedEnrollments(database) })

	fmt.Println("✅ All tables migrated and seeded successfully!")

}

func autoMigrateAndSeed(db *gorm.DB, model interface{}, name string, seedFunc func()) {

	if err := db.AutoMigrate(model); err != nil {
		log.Printf(constant.ErrMigrateDatabase, name, err)
		return
	}

	log.Printf(constant.SuccessMigrateDatabase, name)

	seedFunc()

}
