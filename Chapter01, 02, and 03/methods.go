package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func (p *Person) Grow() {
	p.Age++
}

func (p Person) DoesNotGrow() {
	p.Age++
}

func main() {
	p := Person{"JY", 10}
	p.Grow()
	fmt.Println(p.Age)

	ptr := &p
	ptr.DoesNotGrow()
	fmt.Println(p.Age)
}

// will print
// 11
// 11
