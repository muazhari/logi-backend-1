package configurations

import (
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
	}
	return oneMessageBrokerConfiguration
}
