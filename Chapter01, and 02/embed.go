// This file demonstrates struct embeddeding in Go
package main

import (
	"fmt"
)

// Bird is a sample 'super class' to demonstrate 'inheritance' via embedding
type Bird struct {
	featherLength  int
	classification string
}

// Pigeon is the derived struct
type Pigeon struct {
	Bird
	featherLength  float64
	Name     string
}

func main() {
	p := Pigeon{Name :"Tweety", }
	p.featherLength = 3.14

	fmt.Println(p)
}