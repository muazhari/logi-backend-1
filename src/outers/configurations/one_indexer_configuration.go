package configurations

import "os"

type OneIndexerConfiguration struct {
	Host         string
	ExternalPort string
	InternalPort string
	ExternalUrl  string
	InternalUrl  string
	Username     string
	Password     string
	Index        string
}

func NewOneIndexerConfiguration() *OneIndexerConfiguration {
	oneIndexerConfiguration := &OneIndexerConfiguration{
		Host:         os.Getenv("ELASTICSEARCH_1_HOST"),
		ExternalPort: os.Getenv("ELASTICSEARCH_1_EXTERNAL_PORT"),
		InternalPort: os.Getenv("ELASTICSEARCH_1_INTERNAL_PORT"),
		ExternalUrl:  os.Getenv("ELASTICSEARCH_1_EXTERNAL_URL"),
		InternalUrl:  os.Getenv("ELASTICSEARCH_1_INTERNAL_URL"),
		Username:     os.Getenv("ELASTICSEARCH_1_USERNAME"),
		Password:     os.Getenv("ELASTICSEARCH_1_PASSWORD"),
		Index:        os.Getenv("ELASTICSEARCH_1_INDEX"),
	}
	return oneIndexerConfiguration
}
