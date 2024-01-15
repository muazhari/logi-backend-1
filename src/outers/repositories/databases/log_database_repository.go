package databases

import (
	"context"
	"github.com/muazhari/logi-backend-1/src/inners/models/entities"
	databaseDatastores "github.com/muazhari/logi-backend-1/src/outers/datastores/databases"
	"github.com/muazhari/logi-backend-1/src/outers/utilities"
	"time"
)

type LogDatabaseRepository struct {
	OneDatabaseDatastore *databaseDatastores.OneDatabaseDatastore
}

func NewLogDatabaseRepository(oneDatastore *databaseDatastores.OneDatabaseDatastore) *LogDatabaseRepository {
	logDatabaseRepository := &LogDatabaseRepository{
		OneDatabaseDatastore: oneDatastore,
	}
	return logDatabaseRepository
}

func (logRepository *LogDatabaseRepository) CreateOne(entity *entities.Log) (err error) {
	database := logRepository.OneDatabaseDatastore.Configuration.Database
	databaseQuery := logRepository.OneDatabaseDatastore.Client.Database(database)
	collectionQuery := databaseQuery.Collection("log")

	insertOneCtx, insertOneCtxCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer insertOneCtxCancel()
	entityBsonDocument := utilities.ToBsonDocument(entity)
	_, insertOneQueryErr := collectionQuery.InsertOne(insertOneCtx, entityBsonDocument)
	if insertOneQueryErr != nil {
		err = insertOneQueryErr
	}

	return err
}
