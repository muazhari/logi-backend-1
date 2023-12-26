package managements

import (
	"github.com/muazhari/logi-backend-1/src/outers/repositories/databases"
)

type LogManagement struct {
	LogDatabaseRepository *databases.LogDatabaseRepository
}

func NewLogManagement(logRepository *databases.LogDatabaseRepository) *LogManagement {
	logManagement := &LogManagement{
		LogDatabaseRepository: logRepository,
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
