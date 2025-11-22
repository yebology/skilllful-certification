package main

import (
	"github.com/yebology/skillful-certification/config"
	"github.com/yebology/skillful-certification/cmd/reset/clear"
)

func main() {

	config.LoadEnv()

	config.ConnectDB()

	db := config.GetDB()

	clear.ClearDB(db)

}
