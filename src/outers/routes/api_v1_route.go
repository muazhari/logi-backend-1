package routes

import (
	"github.com/muazhari/logi-backend-1/src/inners/use_cases/managements"
	"github.com/muazhari/logi-backend-1/src/outers/configurations"
	"github.com/muazhari/logi-backend-1/src/outers/controllers/rests"
	databaseDatastores "github.com/muazhari/logi-backend-1/src/outers/datastores/databases"
	indexerDatastores "github.com/muazhari/logi-backend-1/src/outers/datastores/indexers"
	"github.com/muazhari/logi-backend-1/src/outers/datastores/message_brokers"
	databaseRepositories "github.com/muazhari/logi-backend-1/src/outers/repositories/databases"
	indexerRepositories "github.com/muazhari/logi-backend-1/src/outers/repositories/indexers"
	messageBrokerRepositories "github.com/muazhari/logi-backend-1/src/outers/repositories/message_brokers"

	"github.com/gofiber/fiber/v2"
)

type ApiV1Route struct {
	App *fiber.App
}

func NewApiV1Route(app *fiber.App) *ApiV1Route {

	oneDatabaseConfiguration := configurations.NewOneDatabaseConfiguration()
	oneIndexerConfiguration := configurations.NewOneIndexerConfiguration()
	oneMessageBrokerConfiguration := configurations.NewOneMessageBrokerConfiguration()

	oneDatabaseDatastore := databaseDatastores.NewOneDatabaseDatastore(oneDatabaseConfiguration)
	oneDatabaseDatastore.Connect()

	oneIndexerDatastore := indexerDatastores.NewOneIndexerDatastore(oneIndexerConfiguration)
	oneIndexerDatastore.Connect()

	oneMessageBrokerDatastore := message_brokers.NewOneMessageBrokerDatastore(oneMessageBrokerConfiguration)
	oneMessageBrokerDatastore.Connect()

	logDatabaseRepository := databaseRepositories.NewLogDatabaseRepository(oneDatabaseDatastore)
	logIndexerRepository := indexerRepositories.NewLogIndexerRepository(oneIndexerDatastore)
	logMessageBrokerRepository := messageBrokerRepositories.NewLogMessageBrokerRepository(oneMessageBrokerDatastore)

	logManagement := managements.NewLogManagement(logDatabaseRepository, logIndexerRepository, logMessageBrokerRepository)

	apiV1Router := app.Group("/api/v1")

	logController := rests.NewLogController(apiV1Router, logManagement)
	logController.RegisterRoutes()

	apiV1Route := &ApiV1Route{
		App: app,
	}

	return apiV1Route
}
