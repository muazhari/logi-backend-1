package indexers

import (
	indexerDatastores "github.com/muazhari/logi-backend-1/src/outers/datastores/indexers"
)

type LogIndexerRepository struct {
	OneIndexerDatastore *indexerDatastores.OneIndexerDatastore
}

func NewLogIndexerRepository(oneIndexerDatastore *indexerDatastores.OneIndexerDatastore) *LogIndexerRepository {
	logIndexerRepository := &LogIndexerRepository{
		OneIndexerDatastore: oneIndexerDatastore,
	}
	return logIndexerRepository
}
