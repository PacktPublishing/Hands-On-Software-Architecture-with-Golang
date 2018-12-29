package main

import (
	"fmt"
)

// the Node interface
type Node interface {
	Accept(Visitor)
}

type ConcreteNodeX struct{}
func (n ConcreteNodeX) Accept(visitor Visitor) {
	visitor.Visit(n)
}

type ConcreteNodeY struct{}
func (n ConcreteNodeY) Accept(visitor Visitor) {
	// do something NodeY-specific before visiting 
	fmt.Println("ConcreteNodeY being visited !")
	visitor.Visit(n)
}


// the Vistor interface
type Visitor interface {
	Visit(Node)
}

// and an implementation  
type ConcreteVisitor struct{}
func (v ConcreteVisitor) Visit(node Node) {
	fmt.Println("doing something concrete")
            
    // since there is no function overloading..
    // this is one way of checking the concrete node type
	switch node.(type) {
	case ConcreteNodeX:
		fmt.Println("on Node ")
	case ConcreteNodeY:
		fmt.Println("on Node Y")
	}
}


func main() {
	// a simple aggregate
	aggregate := []Node {ConcreteNodeX{}, ConcreteNodeY{},}
	
	// a vistory
	visitor := new(ConcreteVisitor)
	// iterate and visit
	for _, node := range(aggregate){
		node.Accept(visitor)
	}

}
