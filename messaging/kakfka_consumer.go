package main

import (
	"fmt"
	cluster "github.com/bsm/sarama-cluster"
	"log"
	"os"
	"os/signal"
)

func main() {

	// setup config, enable errors and notifications
	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Group.Return.Notifications = true

	// specify Broker co-ordinates and topics of interest
	brokers := []string{"localhost:9092"}
	topics := []string{"topic_a", "topic_b"}

	// connect, and register specifiying the consumer group name
	consumer, err := cluster.NewConsumer(brokers, "my-consumer-group", topics, config)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	// process errors
	go func() {
		for err := range consumer.Errors() {
			log.Printf("Error: %s\n", err.Error())
		}
	}()

	// process notifications
	go func() {
		for ntf := range consumer.Notifications() {
			log.Printf("Rebalanced: %+v\n", ntf)
		}
	}()

	// process messages
	for msg := range consumer.Messages() {
		fmt.Fprintf(os.Stdout, "%s-%d-%d-%s-%s\n",
			msg.Topic,
			msg.Partition,
			msg.Offset,
			msg.Key,
			msg.Value) // <- Actually process message here

		consumer.MarkOffset(msg, "") // Commit offeset for this  message

	}
}
