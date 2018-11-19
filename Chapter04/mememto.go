package main

import (
	"fmt"
)

// Originator
type Originator struct {
	state string
}

func (o *Originator) GetState()  string {
	return o.state
}

func (o *Originator) SetState(state string) {
	fmt.Println("Setting state to " + state)
	o.state = state
}

func (o *Originator) GetMemento() Memento {
	// externalize state to Momemto objct
	return Memento{o.state}
}


func (o *Originator) Restore(memento Memento) {
	// restore state
	o.state = memento.GetState()
}

// Momento
type Memento struct {
	serializedState string
}

func (m *Memento) GetState() string {
	return m.serializedState
}


// caretaker

func Caretaker() {
	
	// assume that A is the original state of the Orginator
	theOriginator := Originator{"A"}
	theOriginator.SetState("A")
	fmt.Println("theOriginator state = ", theOriginator.GetState() )

	// before mutating, get an momemto
	theMomemto := theOriginator.GetMemento()

	// mutate to unclean
	theOriginator.SetState("unclean")
	fmt.Println("theOriginator state = ", theOriginator.GetState() )

	// rollback
	theOriginator.Restore(theMomemto)
	fmt.Println("RESTORED : theOriginator state = ", theOriginator.GetState() )

}

func main() {
	Caretaker()
}