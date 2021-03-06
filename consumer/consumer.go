package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	cluster "github.com/bsm/sarama-cluster"
	"github.com/jjcollinge/sarama-cluster-eventhub-kafka/cfg"
)

func main() {

	config := cfg.CreateSaramaClusterConfig()
	topic := cfg.EventHubName // Rename eventHub name to topic for clarity

	brokers := []string{fmt.Sprintf("%s.servicebus.windows.net:9093", cfg.Namespace)}
	topics := []string{topic}

	/*
		Issue: EventHubs Kafka head appears to be dropping the connection after receiving a JoinGroup call.
		----
		The cluster consumer defined below will run a background loop defined here (vendor/github.com/bsm/sarama-cluster/consumer.go#L313).
		This loop will (amongst other things) attempt to rebalance the kafka subscriptions.
		The rebalance method (vendor/github.com/bsm/sarama-cluster/consumer.go#L559:20) will try to re-join any consumer groups.
		However, the call to JoinGroup (vendor/github.com/Shopify/sarama/broker.go#L315) will fail with the following error:
			"ECONNRESET"
		or
			"connection reset by peer"
		This error is actually raised when trying to read the response from the connection buffer rather than when writting the request.
		Therefore the actual error is raised here: vendor/github.com/Shopify/sarama/broker.go#L695
		and pushed onto an error channel that is then returned from the original JoinGroup sendAndReceive function.
		Therefore it appears the server is dropping the connection on receipt of the JoinGroup request before sending a response.

		If there is an underlying issue with the request i.e. incompatible group procotol - I would expect an error response rather than
		the connection being immediately terminated.
	*/
	consumer, err := cluster.NewConsumer(brokers, cfg.ConsumerGroup, topics, config)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := consumer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	// Trap SIGINT to trigger a shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	// consume errors
	go func() {
		for err := range consumer.Errors() {
			log.Printf("Error: %s\n", err.Error())
		}
	}()

	// consume notifications
	go func() {
		for ntf := range consumer.Notifications() {
			log.Printf("Rebalanced: %+v\n", ntf)
		}
	}()

	// consume messages, watch signals
	for {
		select {
		case msg, ok := <-consumer.Messages():
			if ok {
				fmt.Fprintf(os.Stdout, "%s/%d/%d\t%s\t%s\n", msg.Topic, msg.Partition, msg.Offset, msg.Key, msg.Value)
				consumer.MarkOffset(msg, "") // mark message as processed
			}
		case <-signals:
			return
		}
	}
}
