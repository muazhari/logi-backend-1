package indexers

import (
	"bytes"
	"encoding/json"
	"github.com/muazhari/logi-backend-1/src/inners/models/entities"
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

func (logIndexerRepository *LogIndexerRepository) CreateOne(entity *entities.Log) (err error) {
	marshallOutput, marshallErr := json.Marshal(entity)
	if marshallErr != nil {
		err = marshallErr
	}

	body := bytes.NewReader(marshallOutput)
	index := logIndexerRepository.OneIndexerDatastore.Configuration.Index
	_, indexErr := logIndexerRepository.OneIndexerDatastore.Client.Index(index, body)
	if indexErr != nil {
		err = indexErr
	}

	return err
}
