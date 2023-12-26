package message_brokers

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/muazhari/logi-backend-1/src/outers/configurations"
	"log"
)

type OneIndexerDatastore struct {
	Configuration *configurations.OneIndexerConfiguration
	Client        *elasticsearch.Client
}

func NewOneIndexerDatastore(configuration *configurations.OneIndexerConfiguration) *OneIndexerDatastore {
	return &OneIndexerDatastore{
		Configuration: configuration,
		Client:        nil,
	}
}

func (oneIndexerDatastore *OneIndexerDatastore) Connect() {
	config := elasticsearch.Config{
		Addresses: []string{
			oneIndexerDatastore.Configuration.ExternalUrl,
		},
		Username: oneIndexerDatastore.Configuration.Username,
		Password: oneIndexerDatastore.Configuration.Password,
	}
	client, errConnect := elasticsearch.NewClient(config)
	if errConnect != nil {
		log.Fatal(errConnect)
	}

	oneIndexerDatastore.Client = client
}
