package main

import (
	"fmt"
)

func emitter(till int) <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < till; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func xSquare(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for x := range in {
			out <- x * x
		}
		close(out) // close forward
	}()
	return out
}

func addC(in <-chan int, c int) <-chan int {
	out := make(chan int)
	go func() {
		for x := range in {
			out <- x + c
		}
		close(out) // close forward
	}()
	return out
}

func main() {
	// y = x*x + c
	out := addC(
		xSquare(emitter(3)),
		5)

	for y := range out {
		fmt.Println(y)
	}

	// y = x*x*x*x + c

	out1 := addC(
		xSquare(xSquare(emitter(3))),
		5)

	for y := range out1 {
		fmt.Println(y)
	}

}
