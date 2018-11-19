package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"os"
	"os/signal"
	"strconv"
	"time"
)

// Toy message that we want to send
type Message struct {
	Who          string
	TimeAsString string
}

func main() {

	// Create configuration
	config := sarama.NewConfig()
	// The setting below indicates the  level of  reliability needed
	//  Here we are saying we want all brokers in the ISR to ack
	config.Producer.RequiredAcks = sarama.WaitForAll
	// The total number of times to retry sending a message (default 3).
	config.Producer.Retry.Max = 5

	// you don’t need to give a list of all the brokers, just few seeds which will
	// tell the client about other brokers in the cluster
	brokers := []string{"localhost:9092"}
	asyncProducer, err := sarama.NewAsyncProducer(brokers, config)
	if err != nil {
		// Could not connect
		panic(err)
	}

	defer func() {
		if err := asyncProducer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	// Trap SIGINT to break from the loop and clean up.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	exitProgram := make(chan struct{})

	// Simple while(1) look to send current time.
	var nPublished, nErrors int
	go func() {
		for {
			time.Sleep(5 * time.Second)

			// construct a message
			body := Message{
				Who:          "aProcess",
				TimeAsString: strconv.Itoa(int(time.Now().Unix())),
			}

			// marshall it
			payload, _ := json.Marshal(body)

			msg := &sarama.ProducerMessage{
				Topic: "currentTime",
				Key:   sarama.StringEncoder("aKey"),
				Value: sarama.ByteEncoder(payload),
			}
			select {
			case producer.Input() <- msg:
				nPublished++
				fmt.Println("Produce message")
			case err := <-producer.Errors():
				nErrors++
				fmt.Println("Failed to produce message:", err)
			case <-signals:
				exitProgram <- struct{}{}
			}

			log.Printf("Published: %d; Errors: %d\n", nPublished, nErrors)

		}
	}()

	<-exitProgram // wait here till program gets killed

}
