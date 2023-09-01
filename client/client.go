package client

import "github.com/confluentinc/confluent-kafka-go/kafka"

func ClientInitialization() *kafka.AdminClient {
	KafkaClient, err := kafka.NewAdminClient(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:29092",
	})
	if err != nil {
		panic(err)
	}
	return KafkaClient
}
