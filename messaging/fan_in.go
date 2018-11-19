package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := fanIn(emitter("Source1"), emitter("Source2"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-c) // Display the output of the FanIn channel.
	}

}

// this combines the sources to a Fan-In channel
func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string) // The FanIn channel

	// to avoid blocking, listen to the input channels in separate goroutines
	go func() {
		for {
			c <- <-input1 // Write the message to the FanIn channel, Blocking Call.
		}
	}()

	go func() {
		for {
			c <- <-input2 // Write the message to the FanIn channel, Blocking Call.
		}
	}()

	return c
}

// dummy function for a source
func emitter(name string) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("[%s] says %d", name, i)
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond) // Sleep for some time
		}
	}()

	return c
}
