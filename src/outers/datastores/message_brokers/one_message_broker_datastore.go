package message_brokers

import (
	"context"
	"github.com/muazhari/logi-backend-1/src/outers/configurations"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

type OneMessageBrokerDatastore struct {
	Configuration *configurations.OneMessageBrokerConfiguration
	Client        *kafka.Conn
}

func NewOneMessageBrokerDatastore(oneMessageBrokerDatastore *configurations.OneMessageBrokerConfiguration) *OneMessageBrokerDatastore {
	return &OneMessageBrokerDatastore{
		Configuration: oneMessageBrokerDatastore,
		Client:        nil,
	}
}

func (oneMessageBrokerDatastore *OneMessageBrokerDatastore) Connect() {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, errConnect := kafka.DialLeader(
		ctx,
		"tcp",
		oneMessageBrokerDatastore.Configuration.ExternalUrl,
		oneMessageBrokerDatastore.Configuration.Topic,
		oneMessageBrokerDatastore.Configuration.Partition,
	)
	if errConnect != nil {
		log.Fatal(errConnect)
	}

	oneMessageBrokerDatastore.Client = client
}
