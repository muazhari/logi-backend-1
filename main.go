package main

import (
	"fmt"
	"github.com/muazhari/logi-backend-1/src/outers/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	log.Println("Starting app.")

	errEnvLoad := godotenv.Load()

	if errEnvLoad != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()

	routes.NewApiV1Route(app)

	appHost := os.Getenv("APP_HOST")
	appPort := os.Getenv("APP_PORT")
	errAppListen := app.Listen(fmt.Sprintf("%s:%s", appHost, appPort))
	if errAppListen != nil {
		log.Fatal("Error starting app.")
	}

	log.Println("Stopping app.")
}
