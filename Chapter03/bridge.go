package main

import (
	"fmt"
)



type Reservation struct {
	sellerRef Seller
}

func (r Reservation) Cancel() {
	r.sellerRef.CancelReservation(10) // charge $10 as cancellation feed
}

type PremiumReservation struct {
	Reservation
}

func (r PremiumReservation) Cancel() {
	r.sellerRef.CancelReservation(0) // no charges
}

type Seller interface {
	CancelReservation(charge float64)
}

type InstitutionSeller struct {}

func (s InstitutionSeller) CancelReservation(charge float64) {
	fmt.Println("InstitutionSeller CancelReservation charge =", charge)
}

type SmallScaleSeller struct {}

func (s SmallScaleSeller) CancelReservation(charge float64) {
	fmt.Println("SmallScalSeller CancelReservation charge =", charge)
}



func main() {
	res := Reservation{InstitutionSeller{}}
	res.Cancel()
	premiumRes := PremiumReservation{Reservation{SmallScaleSeller{}}}
	premiumRes.Cancel()
}
