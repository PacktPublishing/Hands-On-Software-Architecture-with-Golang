package main


type Mediator interface {
	AddColleague(colleague Colleague)
}

type MyMediator struct {
	colleagues []Colleague
}

