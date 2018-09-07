package cfg

import (
	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
)

/*
	Define you EventHubs connection details.
*/
const (
	// ConsumerGroup is the name of the consumer group
	ConsumerGroup = "$Default"

	// Private
	connectionStringKey = "$ConnectionString"
)

var (
	// Namespace is the EventHubs namespace
	Namespace = ""
	// EventHubName is the EventHubs name
	EventHubName = ""

	// Private
	connectionString = ""
)

// CreateSaramaClusterConfig creates a new sarama cluster config
func CreateSaramaClusterConfig() *cluster.Config {
	config := cluster.NewConfig()
	config.Version = sarama.V1_1_0_0
	config.Consumer.Return.Errors = true
	config.Group.Return.Notifications = true
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	config.Net.TLS.Enable = true
	config.Net.SASL.Enable = true
	config.Net.SASL.Handshake = true
	config.Net.SASL.User = connectionStringKey
	config.Net.SASL.Password = connectionString
	return config
}

// CreateSaramaConfig creates a new sarama config
func CreateSaramaConfig() *sarama.Config {
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
