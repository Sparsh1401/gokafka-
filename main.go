package main

import (
	"fmt"
	"sync"

	"github.com/Sparsh1401/gokafka/client"
	"github.com/Sparsh1401/gokafka/consumer"
	"github.com/Sparsh1401/gokafka/producer"
)

func main() {
	fmt.Println("Intializing Client")
	Kafka := client.ClientInitialization()
	fmt.Println("Kafka Client Initialized", Kafka)

	/* Topic Creattion part which has to implemented only once to create a Topic otherwise it will show errors */

	// topicConfigs := []kafka.TopicSpecification{
	// 	{
	// 		Topic:             "riders_updates",
	// 		NumPartitions:     2,
	// 		ReplicationFactor: 1,
	// 	},
	// }
	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()
	// topicResults, err := Kafka.CreateTopics(
	// 	ctx,
	// 	topicConfigs,
	// 	kafka.SetAdminOperationTimeout(1000),
	// )
	// if err != nil {
	// 	panic(err)
	// }

	// for _, result := range topicResults {
	// 	if result.Error.Code() != kafka.ErrNoError {
	// 		fmt.Println("Error creating Topic")
	// 	} else {
	// 		fmt.Println("Topic created", result.Topic)
	// 	}
	// }

	var wg sync.WaitGroup

	wg.Add(1)
	go producer.Producer()
	wg.Add(1)
	go consumer.Consumer()

	wg.Wait()
	Kafka.Close()
}
