package indexers

import (
	"bytes"
	"github.com/muazhari/logi-backend-1/src/inners/models/entities"
	indexerDatastores "github.com/muazhari/logi-backend-1/src/outers/datastores/indexers"
	"go.mongodb.org/mongo-driver/bson"
	"log"
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

func (logIndexerRepository *LogIndexerRepository) CreateOne(entity *entities.Log) error {
	marshallOutput, marshallErr := bson.Marshal(entity)
	if marshallErr != nil {
		log.Fatal("Failed to marshall: ", marshallErr)
	}

	body := bytes.NewReader(marshallOutput)
	_, err := logIndexerRepository.OneIndexerDatastore.Client.Index("logs", body)
	if err != nil {
		log.Fatal("Failed to index: ", err)
	}
	return nil
}
