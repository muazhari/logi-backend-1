package main

import (
	"fmt"
	"log"
	"logi-backend-1/src/outers/routes"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	log.Println("Starting app...")

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()

	routes.NewApiV1Route(app)

	appPort := os.Getenv("APP_PORT")
	app.Listen(fmt.Sprintf(":%s", appPort))

	log.Println("Stopping app...")
}
