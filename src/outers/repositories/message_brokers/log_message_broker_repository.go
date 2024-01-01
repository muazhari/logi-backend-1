package indexers

import (
	messageBrokerDatastores "github.com/muazhari/logi-backend-1/src/outers/datastores/message_brokers"
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
