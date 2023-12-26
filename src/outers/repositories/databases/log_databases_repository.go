package databases

import (
	"github.com/muazhari/logi-backend-1/src/inners/models/entities"
	databaseDatastores "github.com/muazhari/logi-backend-1/src/outers/datastores/databases"
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

func (logRepository *LogDatabaseRepository) CreateOne(log *entities.Log) (*entities.Log, error) {
	return log, nil
}

func (logRepository *LogDatabaseRepository) ReadOneById(id string) (*entities.Log, error) {
	return nil, nil
}

func (logRepository *LogDatabaseRepository) ReadMany() ([]*entities.Log, error) {
	return nil, nil
}

func (logRepository *LogDatabaseRepository) UpdateOneById(id string, log *entities.Log) (*entities.Log, error) {
	return log, nil
}

func (logRepository *LogDatabaseRepository) DeleteOneById(id string) (*entities.Log, error) {
	return nil, nil
}
