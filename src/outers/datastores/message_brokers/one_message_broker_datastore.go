package message_brokers

import (
	"github.com/muazhari/logi-backend-1/src/outers/configurations"
	"github.com/segmentio/kafka-go"
)

type OneMessageBrokerDatastore struct {
	Configuration *configurations.OneMessageBrokerConfiguration
	Reader        *kafka.Reader
}

func NewOneMessageBrokerDatastore(oneMessageBrokerDatastore *configurations.OneMessageBrokerConfiguration) *OneMessageBrokerDatastore {
	return &OneMessageBrokerDatastore{
		Configuration: oneMessageBrokerDatastore,
		Reader:        nil,
	}
}

func (oneMessageBrokerDatastore *OneMessageBrokerDatastore) Connect() (err error) {
	configuration := kafka.ReaderConfig{
		Brokers: []string{
			oneMessageBrokerDatastore.Configuration.ExternalUrl,
		},
		Topic:       oneMessageBrokerDatastore.Configuration.Topic,
		Partition:   oneMessageBrokerDatastore.Configuration.Partition,
		MaxBytes:    10e3,
		StartOffset: oneMessageBrokerDatastore.Configuration.StartOffset,
		GroupID:     oneMessageBrokerDatastore.Configuration.GroupID,
	}

	client := kafka.NewReader(configuration)

	oneMessageBrokerDatastore.Reader = client

	return err
}
