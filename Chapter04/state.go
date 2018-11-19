package main

import (
	"fmt"
)

// The State Inteface with the polymorphic methods
type State interface {
	Op1(*Context)
	Op2(*Context)
}

// The Context class
type Context struct {
	state State

}
func (c *Context) Op1() {
	c.state.Op1(c)
}
func (c *Context) Op2() {
	c.state.Op2(c)
}
func (c *Context) SetState(state State) {
	c.state = state
}
func NewContext() *Context{
	c := new(Context)
	c.SetState(new(StateA)) // Initial State
	return c
}


// Concrete States
type StateA struct{}
func (s *StateA) Op1(c *Context) {
	fmt.Println("State A : Op1 ")
}
func (s *StateA) Op2(c *Context) {
	fmt.Println("State A : Op2 ")
	c.SetState(new(StateB)) // <-- State Change!
}

type StateB struct{}
func (s *StateB) Op1(c *Context) {
	fmt.Println("State B : Op1 ")
}
func (s *StateB) Op2(c *Context) {
	fmt.Println("State B : Op2 ")
	c.SetState(new(StateA)) // <-- State Change!
}


func main() {
	context := NewContext()

	// state operations
	context.Op1()
	context.Op2() // <- This changes state to State 2
	context.Op1()
	context.Op2() // <- This changes state  back to State 1

}
