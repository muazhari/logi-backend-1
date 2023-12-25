package routes

import (
	"logi-backend-1/src/inners/use_cases/managements"
	"logi-backend-1/src/outers/configurations"
	"logi-backend-1/src/outers/controllers"
	"logi-backend-1/src/outers/datastores"
	"logi-backend-1/src/outers/repositories"

	"github.com/gofiber/fiber/v2"
)

type ApiV1Route struct {
	app *fiber.App
}

func NewApiV1Route(app *fiber.App) *ApiV1Route {

	oneDatastoreConfiguration := configurations.NewOneDatastoreConfiguration()

	oneDatastore := datastores.NewOneDatastore(oneDatastoreConfiguration)

	logRepository := repositories.NewLogRepository(oneDatastore)

	logManagement := managements.NewLogManagement(logRepository)

	apiV1Router := app.Group("/api/v1")

	logController := controllers.NewLogController(apiV1Router, logManagement)
	logController.RegisterRoutes()

	apiV1Route := &ApiV1Route{
		app: app,
	}

	return apiV1Route
}
