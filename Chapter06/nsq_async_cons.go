// This file describes how we consumers can handle messages asynchronously
package main

import (
	"github.com/nsqio/go-nsq"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// messages received on NSQ are sent on this channel for background processing
var delegateChannel chan *nsq.Message

type MyMessageHandler struct {}
func (h *MyMessageHandler) HandleMessage(m *nsq.Message) error {
	// Disable auto-acknoledgement of messages.
	// This means the actual background consumer from the channel will ack the message.
	// NOTE : If for some reason there is a restart of this instance, messages will be available again after timeout
    m.DisableAutoResponse()
    log.Printf("[nsq handler] received %s\n",string(m.Body))
    log.Println("delegaing..\n")
    delegateChannel <- m
    return nil
}



func main() {

	delegateChannel = make(chan *nsq.Message, 50) // make a buffered channel
	cfg := nsq.NewConfig()
	cfg.MaxInFlight = 1000 //Maximum number of messages to allow in flight (concurrency knob)
	topicName := "my_topic"
	channelName := "my_async_chan"
	cons, err := nsq.NewConsumer(topicName, channelName, cfg)
	if err != nil {
		log.Fatalf(err.Error())
	}

	go func() {
		for m := range delegateChannel {
			err := doSomeWork(m) // some long winded tasks
			if err != nil {
				m.Requeue(-1)
				continue
			}
			m.Finish()
		}
	}()

	// the method below is an alternative to AddHandler to enable concurrent processing
	// the second argument is the number of goroutines to spawn for processing
	cons.AddConcurrentHandlers(&MyMessageHandler{}, 20)

	err = cons.ConnectToNSQD("127.0.0.1:4150")
	if err != nil {
		log.Fatalf(err.Error())
	}

	// Wait for a signal to quit
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	// Stop the consumer
	cons.Stop()
	<-cons.StopChan // wait for cleanup

}

func doSomeWork(m *nsq.Message)  error {
	// some long winded tasks
	log.Printf("received msg %s\n",string(m.Body))
	return nil
}