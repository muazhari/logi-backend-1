package repositories

import (
	"logi-backend-1/src/inners/models/entities"
	"logi-backend-1/src/outers/datastores"
)

type LogRepository struct {
	oneDatastore *datastores.OneDatastore
}

func NewLogRepository(oneDatastore *datastores.OneDatastore) *LogRepository {
	logRepository := &LogRepository{
		oneDatastore: oneDatastore,
	}
	return logRepository
}

func (logRepository *LogRepository) CreateOne(log *entities.Log) (*entities.Log, error) {
	return log, nil
}

func (logRepository *LogRepository) ReadOneById(id string) (*entities.Log, error) {
	return nil, nil
}

func (logRepository *LogRepository) ReadMany() ([]*entities.Log, error) {
	return nil, nil
}

func (logRepository *LogRepository) UpdateOneById(id string, log *entities.Log) (*entities.Log, error) {
	return log, nil
}

func (logRepository *LogRepository) DeleteOneById(id string) (*entities.Log, error) {
	return nil, nil
}
