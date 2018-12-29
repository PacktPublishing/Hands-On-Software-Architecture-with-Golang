package main

import (
	"github.com/nsqio/go-nsq"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// MyMessageHandler is a struct which defines a method to handle messages
type MyMessageHandler struct {
	totalMessages int
}

// This is the message handler callback
func (h *MyMessageHandler) HandleMessage(message *nsq.Message) error {
	h.totalMessages++
	log.Printf("Message no %d received , body : %s \n", h.totalMessages, string(message.Body))
	return nil
}

func main() {
	// Create a consumer
	config := nsq.NewConfig()
	topicName := "my_topic"
	channelName := "my_chan"
	cons, err := nsq.NewConsumer(topicName, channelName, config)
	if err != nil {
		log.Fatal(err)
	}

	// Add the callback for messages
	cons.AddHandler(&MyMessageHandler{})

	// Connect to external NSQD
	err = cons.ConnectToNSQD("127.0.0.1:4150")
	if err != nil {
		log.Panic("Could not connect to NSQD")
	}

	// Wait for a signal to quit
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	// Stop the consumer
	cons.Stop()
	<-cons.StopChan // wait for cleanup

}
