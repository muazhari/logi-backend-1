package utilities

import (
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

func ToBsonDocument(value interface{}) (document *bson.D) {
	marshallOutput, marshallErr := bson.Marshal(value)
	if marshallErr != nil {
		log.Fatal("Failed to marshall: ", marshallErr)
	}

	bsonErr := bson.Unmarshal(marshallOutput, &document)
	if bsonErr != nil {
		log.Fatal("Failed to unmarshall: ", bsonErr)
	}

	return document
}
