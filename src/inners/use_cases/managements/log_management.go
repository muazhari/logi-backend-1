package managements

import (
	"errors"
	"github.com/google/uuid"
	"github.com/muazhari/logi-backend-1/src/inners/models/entities"
	databaseRepositories "github.com/muazhari/logi-backend-1/src/outers/repositories/databases"
	indexerRepositories "github.com/muazhari/logi-backend-1/src/outers/repositories/indexers"
	messageBrokerRepositories "github.com/muazhari/logi-backend-1/src/outers/repositories/message_brokers"
	"github.com/segmentio/kafka-go"
	"time"
)

type LogManagement struct {
	LogDatabaseRepository      *databaseRepositories.LogDatabaseRepository
	LogIndexerRepository       *indexerRepositories.LogIndexerRepository
	LogMessageBrokerRepository *messageBrokerRepositories.LogMessageBrokerRepository
	SettingManagement          *SettingManagement
}

func NewLogManagement(
	logRepository *databaseRepositories.LogDatabaseRepository,
	logIndexerRepository *indexerRepositories.LogIndexerRepository,
	logMessageBrokerRepository *messageBrokerRepositories.LogMessageBrokerRepository,
	settingManagement *SettingManagement,
) *LogManagement {
	logManagement := &LogManagement{
		LogDatabaseRepository:      logRepository,
		LogIndexerRepository:       logIndexerRepository,
		LogMessageBrokerRepository: logMessageBrokerRepository,
		SettingManagement:          settingManagement,
	}
	return logManagement
}

func (logManagement *LogManagement) ConsumeMessage() (err error) {
	consumeMessageErr := logManagement.LogMessageBrokerRepository.ConsumeMessage(
		func(message *kafka.Message) (err error) {
			zone, _ := time.Now().Zone()
			entity := entities.NewLog(
				uuid.New().String(),
				string(message.Value),
				time.Now().Unix(),
				zone,
			)

			readSetting, readSettingErr := logManagement.SettingManagement.ReadOne()
			if readSetting == nil {
				err = errors.New("readSetting is nil")
			}
			if readSettingErr != nil {
				err = readSettingErr
			}

			deleteOlderThanRetainedTimeErr := logManagement.LogIndexerRepository.DeleteManyOlderThanRetainedTime(readSetting.IndexRetainTime)
			if deleteOlderThanRetainedTimeErr != nil {
				err = deleteOlderThanRetainedTimeErr
			}

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
