package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"github.com/muazhari/logi-backend-1/src/outers/routes"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	log.Debugf("Starting app.")

	errEnvLoad := godotenv.Load()
	if errEnvLoad != nil {
		log.Debugf("Error loading .env file: %+v", errEnvLoad)
	}

	app := fiber.New()

	routes.NewApiV1Route(app)

	appHost := os.Getenv("APP_HOST")
	appPort := os.Getenv("APP_PORT")
	errAppListen := app.Listen(fmt.Sprintf("%s:%s", appHost, appPort))
	if errAppListen != nil {
		log.Debugf("Error starting app: %+v", errAppListen)
	}

	log.Debugf("App started.")
}
