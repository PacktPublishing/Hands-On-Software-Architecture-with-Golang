package main

import (
	"fmt"
)

// InterfaceX is the component interface
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

	// if there are children then the component is a composite
	fmt.Println("I'm a composite ")
	for _, child := range c.children {
		child.MethodA()
	}
}

func (c *Composite) AddChild(child InterfaceX) {
	c.children = append(c.children, child)
}




func main() {
	var parent InterfaceX 

	parent =  &Composite{}
	parent.MethodA() // still a leaf

	var child Composite
	parent.AddChild(&child)
	parent.MethodA() // one composite, one  leaf


}