package consumer

import (
	"fmt"
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func Consumer() {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:29092",
		"group.id":          "foo",
		"auto.offset.reset": "smallest"})

	if err != nil {
		log.Fatal(err)
	}
	topic := "riders_updates"
	err = consumer.Subscribe(topic, nil)
	if err != nil {
		log.Fatal(err)
	}

	for {
		ev := consumer.Poll(100)
		switch e := ev.(type) {
		case *kafka.Message:
			fmt.Println("Message recieved from Topic", topic+"Message is:", string(e.Value))
		case kafka.Error:
			fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
		}
	}

	consumer.Close()
}
