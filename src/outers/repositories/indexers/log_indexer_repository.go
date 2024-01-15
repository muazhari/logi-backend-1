package indexers

import (
	"bytes"
	"encoding/json"
	"fmt"
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

	queryBytes := bytes.NewReader(marshallOutput)
	_, indexErr := logIndexerRepository.OneIndexerDatastore.Client.Index("log", queryBytes)
	if indexErr != nil {
		err = indexErr
	}

	return err
}

func (logIndexerRepository *LogIndexerRepository) DeleteManyOlderThanRetainedTime(retainedTime string) (err error) {

	query := fmt.Sprintf(`
	{
		"query": {
			"range": {
				"timestamp": {
					"lt": "now-%s"
				}
			}
		}
	}
	`, retainedTime)

	indexes := []string{"log"}
	queryBytes := bytes.NewReader([]byte(query))
	_, deleteByQueryErr := logIndexerRepository.OneIndexerDatastore.Client.DeleteByQuery(
		indexes,
		queryBytes,
	)
	if deleteByQueryErr != nil {
		err = deleteByQueryErr
	}

	return err
}
