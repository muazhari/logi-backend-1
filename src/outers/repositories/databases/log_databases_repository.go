package databases

import (
	"context"
	"github.com/muazhari/logi-backend-1/src/inners/models/entities"
	databaseDatastores "github.com/muazhari/logi-backend-1/src/outers/datastores/databases"
	"github.com/muazhari/logi-backend-1/src/outers/utilities"
	"log"
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

func (logRepository *LogDatabaseRepository) CreateOne(entity *entities.Log) error {
	insertOneContext, insertOneCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer insertOneCancel()
	bsonDocumentEntity := utilities.ToBsonDocument(entity)
	database := logRepository.OneDatabaseDatastore.Configuration.Database
	result, insertOneErr := logRepository.OneDatabaseDatastore.Client.Database(database).Collection("log").InsertOne(insertOneContext, bsonDocumentEntity)
	if insertOneErr != nil {
		log.Fatal("Failed to insert one entity: ", insertOneErr)
	}
	log.Default().Printf("Inserted a single document: %+v", result)

	return nil
}

func (logRepository *LogDatabaseRepository) ReadOneById(id string) (*entities.Log, error) {
	return nil, nil
}

func (logRepository *LogDatabaseRepository) ReadMany() ([]*entities.Log, error) {
	return nil, nil
}

func (logRepository *LogDatabaseRepository) UpdateOneById(id string, entity *entities.Log) error {
	return nil
}

func (logRepository *LogDatabaseRepository) DeleteOneById(id string) error {
	return nil
}
