package producer

import (
	"fmt"
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func Producer() {
	fmt.Println("Producer producing some message to Topic")
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:29092",
		"client.id":         "dante",
		"acks":              "all",
	})

	if err != nil {
		fmt.Println("Failed to create Producer:", err)
		os.Exit(1)
	}

	fmt.Println("Successfully created Producer:", p)

	delivery_chan := make(chan kafka.Event, 10000)

	topic := "riders_updates"

	err = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: int32(kafka.PartitionAny)},
		Value:          []byte("Hi from kafka"),
	},
		delivery_chan,
	)
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Println("Failed to deliver Message")
				} else {
					fmt.Printf("Successfully produced record to topic %s partition [%d] @ offset %v\n",
						*ev.TopicPartition.Topic, ev.TopicPartition.Partition, ev.TopicPartition.Offset)
				}
			}
		}
	}()
}
