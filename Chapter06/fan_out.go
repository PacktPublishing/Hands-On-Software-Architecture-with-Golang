package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Message struct {
	body string
	key  int
}

func main() {
	evenPipe, oddPipe := fanOut(emitter())
	sink("even", evenPipe)
	sink("odd", oddPipe)

	time.Sleep(10 * time.Second)

}

// this combines the sources to a Fan-In channel
func fanOut(input <-chan Message) (<-chan Message, <-chan Message) {
	even := make(chan Message) // The fan-out channels
	odd := make(chan Message)  // The fan-out channels

	// spawn the fan-out loop
	go func() {
		for {
			msg := <-input
			if msg.key%2 == 0 {
				even <- msg
			} else {
				odd <- msg
			}
		}
	}()

	return even, odd
}

// dummy function for a source
func emitter() <-chan Message {
	c := make(chan Message)

	go func() {
		for i := 0; ; i++ {
			c <- Message{fmt.Sprintf("Message[%d]", i), i}
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond) // Sleep for some time
		}
	}()

	return c
}

func sink(name string, in <-chan Message) {
	go func() {
		for {
			msg := <-in
			fmt.Printf("[%s] says %s\n", name, msg.body)
		}
	}()
}
