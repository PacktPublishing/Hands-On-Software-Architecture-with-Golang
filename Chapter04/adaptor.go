package main

import (
	"fmt"
)


type Adaptee struct{}

func (a *Adaptee) ExistingMethod() {
	fmt.Println("using existing method")
}

type Adapter struct {
	adaptee *Adaptee
}

func NewAdapter() *Adapter {
	return &Adapter{new(Adaptee)}
}

func (a *Adapter) ExpectedMethod() {
	fmt.Println("doing some work")
	a.adaptee.ExistingMethod()
}

func main() {
	adaptor := NewAdapter()
	adaptor.ExpectedMethod()
}