package indexers

import (
	"context"
	log "github.com/gofiber/fiber/v2/log"
	messageBrokerDatastores "github.com/muazhari/logi-backend-1/src/outers/datastores/message_brokers"
	"github.com/segmentio/kafka-go"
	"time"
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

func (logMessageBrokerRepository *LogMessageBrokerRepository) ConsumeMessage(callback func(message *kafka.Message) error) (err error) {
	go func() {
		for {
			readMessageCtx, readMessageCtxCancel := context.WithTimeout(context.Background(), 1*time.Second)
			message, readMessageErr := logMessageBrokerRepository.OneMessageBrokerDatastore.Reader.ReadMessage(readMessageCtx)
			defer readMessageCtxCancel()

			if message.Value == nil {
				continue
			}

			if readMessageErr != nil {
				log.Debugf("readMessageErr: %+v", readMessageErr)
			}
			log.Debugf("message: %+v", message)
			log.Debugf("key: %+v & value: %+v", string(message.Key), string(message.Value))

			callbackErr := callback(&message)
			if callbackErr != nil {
				log.Debugf("callbackErr: %+v", callbackErr)
			}
		}
	}()

	return err
}
