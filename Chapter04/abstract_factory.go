package main

import (
	"fmt"
)

// We have Reservation and Invoice as two generic products
type Reservation interface{}
type Invoice interface{}

type AbstractFactory interface {
	CreateReservation() Reservation
	CreateInvoice() Invoice
}

type HotelFactory struct{}

func (f HotelFactory) CreateReservation() Reservation {
	return new(HotelReservation)
}

func (f HotelFactory) CreateInvoice() Invoice {
	return new(HotelInvoice)
}

type FlightFactory struct{}

func (f FlightFactory) CreateReservation() Reservation {
	return new(FlightReservation)
}

func (f FlightFactory) CreateInvoice() Invoice {
	return new(FlightReservation)
}


type HotelReservation struct{}
type HotelInvoice struct{}
type FlightReservation struct{}
type FlightInvoice struct{}

func  GetFactory(vertical string) AbstractFactory {
	var factory  AbstractFactory
	switch vertical {
	case "flight":
		factory = FlightFactory{}
	case "hotel":
		factory = HotelFactory{}
	}

	return factory
}

func main() {
	hotelFactory := GetFactory("hotel")
	reservation := hotelFactory.CreateReservation()
	invoice := hotelFactory.CreateInvoice()
	
	fmt.Printf("%T\n", reservation)
	fmt.Printf("%T\n", invoice)

}