package configurations

import (
	"os"
)

type OneDatabaseConfiguration struct {
	Host     string
	Port     string
	Database string
	Username string
	Password string
	Url      string
}

func NewOneDatabaseConfiguration() *OneDatabaseConfiguration {
	oneDatabaseConfiguration := &OneDatabaseConfiguration{
		Host:     os.Getenv("MONGO_1_HOST"),
		Port:     os.Getenv("MONGO_1_PORT"),
		Username: os.Getenv("MONGO_1_ROOT_USERNAME"),
		Password: os.Getenv("MONGO_1_PASSWORD"),
		Database: os.Getenv("MONGO_1_DATABASE"),
		Url:      os.Getenv("MONGO_1_URL"),
	}
	return oneDatabaseConfiguration
}
