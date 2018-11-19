package main

import (
	"fmt"
)

type LatLong struct {
	Lat  float64
	Long float64
}

type Animal interface {
	GetLocation() LatLong
	SetLocation(LatLong)
	CanFly() bool
	Speak() string
	GetName() string
}

// The Lion Family
//

type Lion struct {
	name       string
	maneLength int
	location   LatLong
}

func (lion *Lion) GetLocation() LatLong {
	return lion.location
}

func (lion *Lion) SetLocation(loc LatLong) {
	lion.location = loc
}

func (lion *Lion) CanFly() bool {
	return false
}

func (lion *Lion) Speak() string {
	return "roar"
}

func (lion *Lion) GetManeLength() int {
	return lion.maneLength
}

func (lion *Lion) GetName() string {
	return lion.name
}

// The Pigeon Family
//

type Pigeon struct {
	name     string
	location LatLong
}

func (p *Pigeon) GetLocation() LatLong {
	return p.location
}

func (p *Pigeon) SetLocation(loc LatLong) {
	p.location = loc
}

func (p *Pigeon) CanFly() bool {
	return false
}

func (p *Pigeon) Speak() string {
	return "hoot"
}

func (p *Pigeon) GetName() string {
	return p.name
}

// The symphony

func makeThemSing(animals []Animal) {
	for _, animal := range animals {
		fmt.Println(animal.GetName() + " says " + animal.Speak())
	}
}

func main() {
	var myZoo []Animal

	Leo := Lion{
		"Leo",
		10,
		LatLong{10.40, 11.5},
	}
	myZoo = append(myZoo, &Leo)

	Tweety := Pigeon{
		"Tweety",
		LatLong{10.40, 11.5},
	}
	myZoo = append(myZoo, &Tweety)

	makeThemSing(myZoo)


	var aAnimal Animal

	aAnimal =  &Lion{
		"Leo",
		10,
		LatLong{10.40, 11.5},
	}

	fmt.Println(aAnimal)

}
