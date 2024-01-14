package configurations

import (
	"github.com/segmentio/kafka-go"
	"log"
	"os"
	"strconv"
)

type OneMessageBrokerConfiguration struct {
	ExternalHost string
	ExternalPort string
	ExternalUrl  string
	Topic        string
	Partition    int
	StartOffset  int64
	GroupID      string
}

func NewOneMessageBrokerConfiguration() *OneMessageBrokerConfiguration {
	partition, err := strconv.Atoi(os.Getenv("KAFKA_1_PARTITION"))
	if err != nil {
		log.Fatal(err)
	}
	oneMessageBrokerConfiguration := &OneMessageBrokerConfiguration{
		ExternalHost: os.Getenv("KAFKA_1_EXTERNAL_HOST"),
		ExternalPort: os.Getenv("KAFKA_1_EXTERNAL_PORT"),
		Topic:        os.Getenv("KAFKA_1_TOPIC"),
		Partition:    partition,
		ExternalUrl:  os.Getenv("KAFKA_1_EXTERNAL_URL"),
		StartOffset:  kafka.LastOffset,
		GroupID:      os.Getenv("KAFKA_1_GROUP_ID"),
	}
	return oneMessageBrokerConfiguration
}
