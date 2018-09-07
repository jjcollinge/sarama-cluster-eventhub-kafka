# sarama-cluster-eventhub-kafka
This sample demonstrates an issue using the [sarama-cluster](https://github.com/bsm/sarama-cluster) project with Azure EventHubs Kafka API.

## Problem statement

**Issue:** EventHubs Kafka head appears to be dropping the connection after receiving a JoinGroup call.

Please refer to the comment on line 46-63 in `consumer/consumer.go` for details. *Note: this issue is only evident in the consumer as it is only the consumer which uses the sarama-cluster module.*

## Usage
Ensure you have Go installed and have setup your $GOPATH accordingly.


Clone this repository to the path `$GOPATH/src/github.com/jjcollinge/sarama-cluster-eventhubs-kafka` then
follow the guide below:

```
REPO_PATH=$GOPATH/src/github.com/jjcollinge/sarama-cluster-eventhubs-kafka
git clone https://github.com/jjcollinge/sarama-cluster-eventhub-kafka $REPO_PATH
cd $REPO_PATH
go run producer/producer.go
go run consumer/consumer.go
```

## Debugging
Ensure you have installed the [Go extension for vscode](https://github.com/Microsoft/vscode-go).

Run the `Debug producer` or `Debug consumer` configuration.
