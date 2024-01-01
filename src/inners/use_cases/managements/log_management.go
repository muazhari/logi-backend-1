package managements

import (
	databaseRepository "github.com/muazhari/logi-backend-1/src/outers/repositories/databases"
	indexerRepository "github.com/muazhari/logi-backend-1/src/outers/repositories/indexers"
	messageBrokerRepository "github.com/muazhari/logi-backend-1/src/outers/repositories/message_brokers"
)

type LogManagement struct {
	LogDatabaseRepository      *databaseRepository.LogDatabaseRepository
	LogIndexerRepository       *indexerRepository.LogIndexerRepository
	LogMessageBrokerRepository *messageBrokerRepository.LogMessageBrokerRepository
}

func NewLogManagement(
	logRepository *databaseRepository.LogDatabaseRepository,
	logIndexerRepository *indexerRepository.LogIndexerRepository,
	logMessageBrokerRepository *messageBrokerRepository.LogMessageBrokerRepository,
) *LogManagement {
	logManagement := &LogManagement{
		LogDatabaseRepository:      logRepository,
		LogIndexerRepository:       logIndexerRepository,
		LogMessageBrokerRepository: logMessageBrokerRepository,
	}
	return logManagement
}

func (logManagement *LogManagement) CreateOne() error {
	return nil
}

func (logManagement *LogManagement) ReadOneById() error {
	return nil
}

func (logManagement *LogManagement) ReadMany() error {
	return nil
}

func (logManagement *LogManagement) UpdateOneById() error {
	return nil
}

func (logManagement *LogManagement) DeleteOneById() error {
	return nil
}
