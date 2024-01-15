package seeders

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/muazhari/logi-backend-1/src/inners/models/value_objects"
	databaseRepositories "github.com/muazhari/logi-backend-1/src/outers/repositories/databases"
	indexerRepositories "github.com/muazhari/logi-backend-1/src/outers/repositories/indexers"
	messageBrokerRepositories "github.com/muazhari/logi-backend-1/src/outers/repositories/message_brokers"
)

type OneSeeder struct {
	LogDatabaseRepository      *databaseRepositories.LogDatabaseRepository
	LogIndexerRepository       *indexerRepositories.LogIndexerRepository
	LogMessageBrokerRepository *messageBrokerRepositories.LogMessageBrokerRepository
	SettingDatabaseRepository  *databaseRepositories.SettingDatabaseRepository
}

func NewOneSeeder(
	logDatabaseRepository *databaseRepositories.LogDatabaseRepository,
	logIndexerRepository *indexerRepositories.LogIndexerRepository,
	logMessageBrokerRepository *messageBrokerRepositories.LogMessageBrokerRepository,
	SettingDatabaseRepository *databaseRepositories.SettingDatabaseRepository,
) *OneSeeder {
	oneSeeder := &OneSeeder{
		LogDatabaseRepository:      logDatabaseRepository,
		LogIndexerRepository:       logIndexerRepository,
		LogMessageBrokerRepository: logMessageBrokerRepository,
		SettingDatabaseRepository:  SettingDatabaseRepository,
	}
	return oneSeeder
}

func (oneSeeder *OneSeeder) Seed() (err error) {
	log.Debugf("Seeding started.")
	seedSettingErr := oneSeeder.SeedSetting()
	if seedSettingErr != nil {
		err = seedSettingErr
	}
	log.Debugf("Seeding finished.")
	return err
}

func (oneSeeder *OneSeeder) SeedSetting() (err error) {
	setting := value_objects.NewSetting(
		"10m",
	)

	createOneErr := oneSeeder.SettingDatabaseRepository.CreateOne(setting)
	if createOneErr != nil {
		err = createOneErr
	}

	return err
}
