package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/yebology/skillful-certification/config"
	"github.com/yebology/skillful-certification/router"
)

func main() {

	app := fiber.New()

	config.LoadEnv()

	config.ConnectDB()
	config.CheckConnection()

	app.Use(cors.New(cors.Config{
		AllowHeaders: "*",
		AllowOrigins: "*",
	}))

	router.SetUp(app)

	log.Fatal(app.Listen(":8080"))

}