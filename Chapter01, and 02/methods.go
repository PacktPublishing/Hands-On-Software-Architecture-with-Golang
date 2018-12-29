// This file demonstrates methods and pointer v/s value receiver
package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

// Grow method has a pointer receiver. This is Pass-By-Reference
func (p *Person) Grow() {
	p.Age++
}

// DoesNotGrow method has a value receiver. This is Pass-By-Value. Age will nto be modified here
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
