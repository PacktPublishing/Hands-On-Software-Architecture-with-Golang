package main

import (
	"fmt"
	"math/rand"
	"time"
)

// this combines the sources to a Fan-In channel
func fanInSelect(input1, input2 <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		for {
			select {
			case in := <-input1:
				out <- in
			case in := <-input2:
				out <- in
			}
		}
	}()
	return out
}

func main() {
	c := fanInSelect(emitter("Source1"), emitter("Source2"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-c) // Display the output of the FanIn channel.
	}

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
