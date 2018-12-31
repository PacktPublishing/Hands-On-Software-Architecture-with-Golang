package main

import (
	"fmt"
	"time"
)

type Request struct {
	someArg string
	replyTo chan<- Response
}

type Response struct {
	reply string
}

func responder(c <-chan Request) {
	for request := range c {
		var resp Response
		resp.reply = "reply-to-" + request.someArg
		request.replyTo <- resp
	}
}

func requestor(c chan<- Request) {
	myChannel := make(chan Response)
	for i := 0; i < 5; i++ {
		c <- Request{fmt.Sprintf("message%d", i), myChannel}
		resp := <-myChannel
		fmt.Printf("request %d , response %s\n", i, resp.reply)
	}

	// cleanup after my work is done
	close(myChannel)
}

func main() {

	requestChannel := make(chan Request)
	go responder(requestChannel)
	go requestor(requestChannel)

	time.Sleep(time.Second * 10)

}
