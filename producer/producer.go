package main

import (
	"fmt"

	"github.com/Shopify/sarama"
)

/*
	Define you EventHubs connection details.
*/
const (
	namespace           = "my-namespace"
	eventHubName        = "my-eventhub"
	connectionStringKey = "$ConnectionString"
	connectionString    = "my-connection-string"
	consumerGroup       = "$Default"
)

func createSaramaConfig() *sarama.Config {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	config.Version = sarama.V1_1_0_0
	config.Net.TLS.Enable = true
	config.Net.SASL.Enable = true
	config.Net.SASL.Handshake = true
	config.Net.SASL.User = connectionStringKey
	config.Net.SASL.Password = connectionString
	return config
}

func main() {

	config := createSaramaConfig()
	topic := eventHubName // Rename eventHub name to topic for clarity

	brokers := []string{fmt.Sprintf("%s.servicebus.windows.net:9093", namespace)}

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
