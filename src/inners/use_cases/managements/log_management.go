package managements

import (
	"logi-backend-1/src/outers/repositories"
)

type LogManagement struct {
	logRepository *repositories.LogRepository
}

func NewLogManagement(logRepository *repositories.LogRepository) *LogManagement {
	logManagement := &LogManagement{
		logRepository: logRepository,
	}
	return logManagement
}
