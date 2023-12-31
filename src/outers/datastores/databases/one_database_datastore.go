package databases

import (
	"github.com/muazhari/logi-backend-1/src/outers/configurations"
	"log"
	"time"

	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type OneDatabaseDatastore struct {
	Configuration *configurations.OneDatabaseConfiguration
	Client        *mongo.Client
}

func NewOneDatabaseDatastore(oneDatastoreConfiguration *configurations.OneDatabaseConfiguration) *OneDatabaseDatastore {
	oneDatabaseDatastore := &OneDatabaseDatastore{
		Configuration: oneDatastoreConfiguration,
		Client:        nil,
	}
	return oneDatabaseDatastore
}

func (oneDatastore *OneDatabaseDatastore) Connect() {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(oneDatastore.Configuration.Url))
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	oneDatastore.Client = client
}
