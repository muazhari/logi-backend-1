package managements

import (
	"github.com/google/uuid"
	"github.com/muazhari/logi-backend-1/src/inners/models/entities"
	databaseRepository "github.com/muazhari/logi-backend-1/src/outers/repositories/databases"
	indexerRepository "github.com/muazhari/logi-backend-1/src/outers/repositories/indexers"
	messageBrokerRepository "github.com/muazhari/logi-backend-1/src/outers/repositories/message_brokers"
	"github.com/segmentio/kafka-go"
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

func (logManagement *LogManagement) ConsumeMessage() (err error) {

	consumeMessageErr := logManagement.LogMessageBrokerRepository.ConsumeMessage(
		func(message *kafka.Message) (err error) {
			entity := entities.NewLog(
				uuid.New().String(),
				string(message.Value),
			)

			createOneToIndexerErr := logManagement.CreateOneToIndexer(entity)
			if createOneToIndexerErr != nil {
				err = createOneToIndexerErr
			}

			createOneToDatabaseErr := logManagement.CreateOneToDatabase(entity)
			if createOneToDatabaseErr != nil {
				err = createOneToDatabaseErr
			}

			return err
		},
	)
	if consumeMessageErr != nil {
		err = consumeMessageErr
	}

	return err
}

func (logManagement *LogManagement) CreateOneToIndexer(entity *entities.Log) (err error) {
	indexerCreateOneErr := logManagement.LogIndexerRepository.CreateOne(entity)
	if indexerCreateOneErr != nil {
		err = indexerCreateOneErr
	}
	return err
}

func (logManagement *LogManagement) CreateOneToDatabase(entity *entities.Log) (err error) {
	databaseCreateOneErr := logManagement.LogDatabaseRepository.CreateOne(entity)
	if databaseCreateOneErr != nil {
		err = databaseCreateOneErr
	}
	return err
}
