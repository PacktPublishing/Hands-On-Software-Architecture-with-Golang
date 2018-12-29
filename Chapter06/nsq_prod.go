package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"log"
)

func main() {

	// Connect
	pCfg := nsq.NewConfig()
	producer, err := nsq.NewProducer("127.0.0.1:4160", pCfg)
	if err != nil {
		log.Fatalf("failed creating producer %s", err)
	}

	// Publish Async
	destinationTopic := "my_topic"
	responseChan := make(chan *ProducerTransaction)
	err = producer.PublishAsync(destinationTopic, []byte("a_message"), responseChan, "some_args")

	// Check for status
	// Done here  inline just for showcase
	status := <-responseChan
	if status.Error != nil {
		log.Printf("Error received %s \n", status.Error.Error())
	} else {
		log.Printf("Success Arg received : %s \n", status.Args[0].(string)) // should be some_args
	}

}
