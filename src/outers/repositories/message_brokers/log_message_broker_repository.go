package indexers

import (
	messageBrokerDatastores "github.com/muazhari/logi-backend-1/src/outers/datastores/message_brokers"
)

type LogMessageBrokerRepository struct {
	OneMessageBrokerDatastore *messageBrokerDatastores.OneMessageBrokerDatastore
}

func NewLogIndexerRepository(oneMessageBrokerDatastore *messageBrokerDatastores.OneMessageBrokerDatastore) *LogMessageBrokerRepository {
	logIndexerRepository := &LogMessageBrokerRepository{
		OneMessageBrokerDatastore: oneMessageBrokerDatastore,
	}
	return logIndexerRepository
}
