// this file demonstrates the bridge design pattern
package main

import (
	"fmt"
)


// Reservation is the main abstraction
//
// the abstraction here is a struct not an interface , since in Go you cant have abstract structs/interfaces ,
// where one can store reference to the Seller implementation
type Reservation struct {
	// sellerRef is the reference to the implementation
	sellerRef Seller
}

func (r Reservation) Cancel() {
	r.sellerRef.CancelReservation(10) // charge $10 as cancellation feed
}

// PremiumReservation is a specialization of the interface
type PremiumReservation struct {
	Reservation
}

func (r PremiumReservation) Cancel() {
	r.sellerRef.CancelReservation(0) // no charges
}

// This is the interface for all Sellers
type Seller interface {
	CancelReservation(charge float64)
}

// InstitutionSeller is an specialization of the implementation
type InstitutionSeller struct {}

func (s InstitutionSeller) CancelReservation(charge float64) {
	fmt.Println("InstitutionSeller CancelReservation charge =", charge)
}


// SmallScaleSeller is an specialization of the implementation
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
