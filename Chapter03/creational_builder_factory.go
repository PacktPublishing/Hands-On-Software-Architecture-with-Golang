package main

import (
	"fmt"
)

type Reservation interface {
	GetReservationDate() string
	SetReservationDate(string)
}

type HotelReservationImpl struct {
	reservationDate string
}

type FlightReservationImpl struct {
	reservationDate string
}

func (r HotelReservationImpl) GetReservationDate() string {
	return r.reservationDate
}

func (r HotelReservationImpl) SetReservationDate(date string) {
	r.reservationDate = date
}

func (r FlightReservationImpl) GetReservationDate() string {
	return r.reservationDate
}

func (r FlightReservationImpl) SetReservationDate(date string) {
	r.reservationDate = date
}

func NewReservation(vertical, reservationDate string) Reservation {
	switch vertical {
	case "flight":
		return FlightReservationImpl{reservationDate}
	case "hotel":
		return HotelReservationImpl{reservationDate}
	default:
		return nil
	}

}

type ReservationBuilder interface {
	Vertical(string) ReservationBuilder
	ReservationDate(string) ReservationBuilder
	Build() Reservation
}

type reservationBuilder struct {
	vertical string
	rdate    string
}

func (r *reservationBuilder) Vertical(v string) ReservationBuilder {
	r.vertical = v
	return r
}

func (r *reservationBuilder) ReservationDate(date string) ReservationBuilder {
	r.rdate = date
	return r
}

func (r *reservationBuilder) Build() Reservation {

	var builtReservation Reservation

	switch r.vertical {
	case "flight":
		builtReservation = FlightReservationImpl{r.rdate}
	case "hotel":
		builtReservation = HotelReservationImpl{r.rdate}
	}

	return builtReservation
}

func NewReservationBuilder() ReservationBuilder {
	return &reservationBuilder{}
}

func main() {
	hotelReservation := NewReservation("hotel", "20180101")
	fmt.Println(hotelReservation.GetReservationDate())

	builder := NewReservationBuilder()
	r := builder.Vertical("hotel").ReservationDate("20180101").Build()
	fmt.Println(r.GetReservationDate())

}
