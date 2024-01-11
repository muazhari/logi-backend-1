package managements

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/muazhari/logi-backend-1/src/inners/models/entities"
	databaseRepository "github.com/muazhari/logi-backend-1/src/outers/repositories/databases"
	indexerRepository "github.com/muazhari/logi-backend-1/src/outers/repositories/indexers"
	messageBrokerRepository "github.com/muazhari/logi-backend-1/src/outers/repositories/message_brokers"
	"github.com/segmentio/kafka-go"
	"log"
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

func (logManagement *LogManagement) ConsumeMessage() error {

	consumeMessageErr := logManagement.LogMessageBrokerRepository.ConsumeMessage(
		func(message *kafka.Message) error {
			entity := entities.NewLog(
				uuid.New().String(),
				fmt.Sprintf("%+v", message),
			)

			createOneToIndexerErr := logManagement.CreateOneToIndexer(entity)
			if createOneToIndexerErr != nil {
				log.Fatal("Failed to create one entity to the indexer: ", createOneToIndexerErr)
			}
			createOneToDatabaseErr := logManagement.CreateOneToDatabase(entity)
			if createOneToDatabaseErr != nil {
				log.Fatal("Failed to create one entity to the database: ", createOneToDatabaseErr)
			}

			return nil
		},
	)
	if consumeMessageErr != nil {
		log.Fatal("Failed to consume message: ", consumeMessageErr)
	}

	return nil
}

func (logManagement *LogManagement) CreateOneToIndexer(entity *entities.Log) error {
	indexerCreateOneErr := logManagement.LogIndexerRepository.CreateOne(entity)
	if indexerCreateOneErr != nil {
		log.Fatal("Failed to create one entity to the indexer: ", indexerCreateOneErr)
	}
	return nil
}

func (logManagement *LogManagement) CreateOneToDatabase(entity *entities.Log) error {
	databaseCreateOneErr := logManagement.LogDatabaseRepository.CreateOne(entity)
	if databaseCreateOneErr != nil {
		log.Fatal("Failed to create one entity to the database: ", databaseCreateOneErr)
	}
	return nil
}
