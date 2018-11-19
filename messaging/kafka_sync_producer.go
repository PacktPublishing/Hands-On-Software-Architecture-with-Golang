package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func main() {

	// Config
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Errors = true  // For sync producer this needs to be true
	config.Producer.Return.Success = true // For sync producer this needs to be true

	// Connect to a Kafka broker running locally
	brokers := []string{"localhost:9092"}
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		panic(err)
	}

	// cleanup
	defer func() {
		if err := producer.Close(); err != nil {
			panic(err)
		}
	}()

	msg := &sarama.ProducerMessage{
		Topic: "currentTime",
		Value: sarama.StringEncoder(strconv.Itoa(int(time.Now().Unix()))),
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		fmt.Printf("FAILED to publish message: %s\n", err)
	} else {
		fmt.Printf("message sent | partition(%d)/offset(%d)\n", partition, offset)
	}
}
