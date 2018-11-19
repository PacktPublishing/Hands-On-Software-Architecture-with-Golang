package main

import (
	"github.com/nsqio/go-nsq"
	"log"
)

type MyMessageHandler struct {
	totalMessages int
}

func (h *MyMessageHandler) HandleMessage(message *Message) error {
	h.totalMessages++
	log.Printf("Message no %d received , body : %s \n", h.totalMessages, string(message.Body))
}

func main() {
	config := NewConfig()

	topicName := "my_topic"
	channelName := "my_chan"
	cons, err := NewConsumer(topicName, channelName, config)
	if err != nil {
		log.Fatal(err)
	}

	cons.AddHandler(&MyMessageHandler{})

	// wait for a signal to quit
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	// Stop the consumer
	cons.Stop()
	<-cons.StopChan // wait for cleanup

}
