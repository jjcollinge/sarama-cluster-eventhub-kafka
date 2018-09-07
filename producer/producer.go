package main

import (
	"fmt"

	"github.com/Shopify/sarama"
	"github.com/jjcollinge/sarama-cluster-eventhub-kafka/cfg"
)

func main() {

	config := cfg.CreateSaramaConfig()
	topic := cfg.EventHubName // Rename eventHub name to topic for clarity

	brokers := []string{fmt.Sprintf("%s.servicebus.windows.net:9093", cfg.Namespace)}

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		// Should not reach here
		panic(err)
	}

	defer func() {
		if err := producer.Close(); err != nil {
			// Should not reach here
			panic(err)
		}
	}()

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder("Something Cool"),
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", topic, partition, offset)
}
