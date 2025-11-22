package config

import (
	"log"

	"github.com/yebology/skillful-certification/constant"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDB() {

	dbUrl := LoadEnvConfig("DIRECT_URL")

	var err error
	db, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  dbUrl,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: false,
	})

	if err != nil {
		log.Fatalln(constant.ErrConnectingDatabase)
	}

}

func CheckConnection() {

	sqlDb, err := db.DB()
	if err != nil {
		log.Fatalf(string(constant.ErrGetSQLInstance))
	}

	err = sqlDb.Ping()
	if err != nil {
		log.Fatalln(constant.ErrPingDatabase)
	}

	log.Println(constant.SuccessConnectDatabase)

}

func GetDB() *gorm.DB {

	return db

}
