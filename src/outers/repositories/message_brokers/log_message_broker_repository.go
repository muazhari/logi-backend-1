package indexers

import (
	messageBrokerDatastores "github.com/muazhari/logi-backend-1/src/outers/datastores/message_brokers"
	"github.com/segmentio/kafka-go"
	"log"
)

type LogMessageBrokerRepository struct {
	OneMessageBrokerDatastore *messageBrokerDatastores.OneMessageBrokerDatastore
}

func NewLogMessageBrokerRepository(oneMessageBrokerDatastore *messageBrokerDatastores.OneMessageBrokerDatastore) *LogMessageBrokerRepository {
	logMessageBrokerRepository := &LogMessageBrokerRepository{
		OneMessageBrokerDatastore: oneMessageBrokerDatastore,
	}
	return logMessageBrokerRepository
}

func (logMessageBrokerRepository *LogMessageBrokerRepository) ConsumeMessage(callback func(message *kafka.Message) error) error {
	go func() {
		for {
			message, readMessageErr := logMessageBrokerRepository.OneMessageBrokerDatastore.Client.ReadMessage(10e3)
			if readMessageErr != nil {
				log.Fatal("Failed to read message: ", readMessageErr)
			}
			log.Default().Printf("Consumed message: %+v", message)

			callbackErr := callback(&message)
			if callbackErr != nil {
				log.Fatal("Failed to execute callback: ", callbackErr)
			}
		}
	}()

	return nil
}
