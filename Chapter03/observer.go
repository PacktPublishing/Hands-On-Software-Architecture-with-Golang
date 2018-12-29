package main


import (
	"fmt"
)

// The Subject

type Subject struct {
	observers []Observer
	state string
}

func (s *Subject) Attach(observer Observer) {
	s.observers = append(s.observers, observer)
}


func (s *Subject) SetState(newState string) {
	s.state = newState
	for _,o := range(s.observers) {
		o.Update()
	}
}

func (s *Subject) GetState()  string {
	return s.state 
}

// The Observer Inteface

type Observer interface {
	Update()
}

// Concrete Observer A

type ConcreteObserverA struct {
	model *Subject
	viewState string
}

func (ca *ConcreteObserverA) Update()   {
	ca.viewState = ca.model.GetState()
	fmt.Println("ConcreteObserverA : updated view state to ", ca.viewState)
}

func (ca *ConcreteObserverA) SetModel(s *Subject)   {
	ca.model = s
}


func main() {
	// create Subject
	s := Subject{}

	// create concrete observer
	ca := &ConcreteObserverA{}
	ca.SetModel(&s) // set Model

	// Attach the observer
	s.Attach(ca) 

	s.SetState("s1")
 
}

