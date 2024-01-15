package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/muazhari/logi-backend-1/src/inners/use_cases/managements"
	"github.com/muazhari/logi-backend-1/src/outers/configurations"
	"github.com/muazhari/logi-backend-1/src/outers/controllers/rests"
	databaseDatastores "github.com/muazhari/logi-backend-1/src/outers/datastores/databases"
	indexerDatastores "github.com/muazhari/logi-backend-1/src/outers/datastores/indexers"
	messageBrokerDatastores "github.com/muazhari/logi-backend-1/src/outers/datastores/message_brokers"
	databaseRepositories "github.com/muazhari/logi-backend-1/src/outers/repositories/databases"
	indexerRepositories "github.com/muazhari/logi-backend-1/src/outers/repositories/indexers"
	messageBrokerRepositories "github.com/muazhari/logi-backend-1/src/outers/repositories/message_brokers"
	"github.com/muazhari/logi-backend-1/src/outers/seeders"
)

type ApiV1Route struct {
	App *fiber.App
}

func NewApiV1Route(app *fiber.App) *ApiV1Route {
	oneDatabaseConfiguration := configurations.NewOneDatabaseConfiguration()
	oneIndexerConfiguration := configurations.NewOneIndexerConfiguration()
	oneMessageBrokerConfiguration := configurations.NewOneMessageBrokerConfiguration()

	oneDatabaseDatastore := databaseDatastores.NewOneDatabaseDatastore(oneDatabaseConfiguration)
	oneDatabaseDatastoreConnectErr := oneDatabaseDatastore.Connect()
	if oneDatabaseDatastoreConnectErr != nil {
		log.Debugf("oneDatabaseDatastoreConnectErr: %+v", oneDatabaseDatastoreConnectErr)
	}

	oneIndexerDatastore := indexerDatastores.NewOneIndexerDatastore(oneIndexerConfiguration)
	oneIndexerDatastoreConnectErr := oneIndexerDatastore.Connect()
	if oneIndexerDatastoreConnectErr != nil {
		log.Debugf("oneIndexerDatastoreConnectErr: %+v", oneIndexerDatastoreConnectErr)
	}

	oneMessageBrokerDatastore := messageBrokerDatastores.NewOneMessageBrokerDatastore(oneMessageBrokerConfiguration)
	oneMessageBrokerDatastoreConnectErr := oneMessageBrokerDatastore.Connect()
	if oneMessageBrokerDatastoreConnectErr != nil {
		log.Debugf("oneMessageBrokerDatastoreConnectErr: %+v", oneMessageBrokerDatastoreConnectErr)
	}

	logDatabaseRepository := databaseRepositories.NewLogDatabaseRepository(oneDatabaseDatastore)
	logIndexerRepository := indexerRepositories.NewLogIndexerRepository(oneIndexerDatastore)
	logMessageBrokerRepository := messageBrokerRepositories.NewLogMessageBrokerRepository(oneMessageBrokerDatastore)

	settingDatabaseRepository := databaseRepositories.NewSettingDatabaseRepository(oneDatabaseDatastore)

	oneSeeder := seeders.NewOneSeeder(
		logDatabaseRepository,
		logIndexerRepository,
		logMessageBrokerRepository,
		settingDatabaseRepository,
	)

	oneSeederErr := oneSeeder.Seed()
	if oneSeederErr != nil {
		log.Debugf("oneSeederErr: %+v", oneSeederErr)
	}

	settingManagement := managements.NewSettingManagement(settingDatabaseRepository)

	logManagement := managements.NewLogManagement(
		logDatabaseRepository,
		logIndexerRepository,
		logMessageBrokerRepository,
		settingManagement,
	)

	logManagementConsumeMessageErr := logManagement.ConsumeMessage()
	if logManagementConsumeMessageErr != nil {
		log.Debugf("logManagementConsumeMessageErr: %+v", logManagementConsumeMessageErr)
	}

	apiV1Router := app.Group("/api/v1")

	logController := rests.NewLogController(apiV1Router, logManagement)
	logController.RegisterRoutes()

	apiV1Route := &ApiV1Route{
		App: app,
	}

	return apiV1Route
}
