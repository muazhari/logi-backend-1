package datastores

import (
	"log"
	"logi-backend-1/src/outers/configurations"
	"time"

	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type OneDatastore struct {
	client *mongo.Client
}

func NewOneDatastore(oneDatastoreConfiguration *configurations.OneDatastoreConfiguration) *OneDatastore {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(oneDatastoreConfiguration.Url))
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	oneDatastore := &OneDatastore{
		client: client,
	}
	return oneDatastore
}
