# sarama-cluster-eventhub-kafka
This sample demonstrates an issue using the [sarama-cluster](https://github.com/bsm/sarama-cluster) project with Azure EventHubs Kafka API.

## Problem statement

**Issue:** EventHubs Kafka head appears to be dropping the connection after receiving a JoinGroup call.

Please refer to the comment on line 46-63 in `consumer/consumer.go` for details. *Note: this issue is only evident in the consumer as it is only the consumer which uses the sarama-cluster module.*

## Usage
Ensure you have Go installed and have setup your $GOPATH accordingly.


1. Clone this repository to the path `$GOPATH/src/github.com/jjcollinge/sarama-cluster-eventhubs-kafka`

```
export REPO_PATH=$GOPATH/src/github.com/jjcollinge/sarama-cluster-eventhubs-kafka
git clone https://github.com/jjcollinge/sarama-cluster-eventhub-kafka $REPO_PATH
```

2. Add you EventHubs details to `consumer/consumer.go` and `producer/producer.go`.

3. Run the producer then the consumer

```
cd $REPO_PATH
go run producer/producer.go
go run consumer/consumer.go
```

4. Observe "connection reset" errors on the consumer during rebalance.

## Debugging
Ensure you have installed the [Go extension for vscode](https://github.com/Microsoft/vscode-go).

Run the `Debug producer` or `Debug consumer` configuration.
