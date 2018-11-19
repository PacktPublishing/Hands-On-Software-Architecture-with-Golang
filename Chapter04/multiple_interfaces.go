package main

import (
	"fmt"
)

type Reservation interface {
	GetReservationDate() string
	CalculateCancellationFee() float64
}

type HotelReservation interface {
	Reservation
	ChangeType()
}

type FlightReservation interface {
	Reservation
	AddExtraLuggageAllowance(peices int)
}

type HotelReservationImpl struct {
	reservationDate string
}

func (r HotelReservationImpl) GetReservationDate() string {
	return r.reservationDate
}

func (r HotelReservationImpl) CalculateCancellationFee() float64 {
	return 1.0 // flat :P
}

type FlightReservationImpl struct {
	reservationDate string
	luggageAllowed  int
}

func (r FlightReservationImpl) AddExtraLuggageAllowance(peices int) {
	r.luggageAllowed = peices
}

func (r FlightReservationImpl) CalculateCancellationFee() float64 {
	return 2.0 // flat but slight more than hotels :P
}

func (r FlightReservationImpl) GetReservationDate() string {
	// this might look repetitive, but the idea is to provide freedom for the
	// derived classes to flux independently of each other
	return r.reservationDate
}

type Trip struct {
	reservations []Reservation
}

func (t *Trip) CalculateCancellationFee() float64 {
	total := 0.0

	for _, r := range t.reservations {
		total += r.CalculateCancellationFee()
	}

	return total
}

func (t *Trip) AddReservation(r Reservation) {
	t.reservations = append(t.reservations, r)
}

func main() {
	var (
		h HotelReservationImpl
		f FlightReservationImpl
		t Trip
	)

	fmt.Println(f.CalculateCancellationFee())
	fmt.Println(h.CalculateCancellationFee())

	t.AddReservation(h)
	t.AddReservation(f)
	fmt.Println(t.CalculateCancellationFee())

}
