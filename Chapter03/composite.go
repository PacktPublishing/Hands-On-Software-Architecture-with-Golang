package main

import (
	"fmt"
	"time"
	"math"
)

type InterfaceX interface {
	MethodA()
	AddChild(InterfaceX)
}

type Composite struct{
	children []InterfaceX
}

func (c *Composite) MethodA() {
	if len(c.children) == 0 {
		fmt.Println("I'm a leaf ")
		return
	}

	fmt.Println("I'm a composite ")
	for _, child := range c.children {
		child.MethodA()
	}
}

func (c *Composite) AddChild(child InterfaceX) {
	c.children = append(c.children, child)
}


type Function func(float64) float64

func ProfileDecorator(fn Function) Function {
	return func(params float64) float64 {
		start := time.Now()
		result := fn(params)
		elapsed := time.Now().Sub(start)
		fmt.Println("Funtion completed with time : ", elapsed)

        return result
	}
}


func SquareRoot(n float64) float64 {
    return math.Sqrt(n)
}

func main() {
	var parent InterfaceX 

	parent =  &Composite{}
	parent.MethodA() // still a leaf

	var child Composite
	parent.AddChild(&child)
	parent.MethodA() // one composite, one  leaf

	decoratedSqaureRoot := ProfileDecorator(SquareRoot)
	fmt.Println(decoratedSqaureRoot(16))

}
