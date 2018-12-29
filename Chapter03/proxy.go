package  main

import (
	"fmt"
)


// Proxy
type HotelBoutiqueProxy struct {
	subject *HotelBoutique
}

func (p *HotelBoutiqueProxy) Book() {
	if p.subject == nil {
		p.subject = new(HotelBoutique)
	}
	fmt.Println("Proxy Delegating Booking call")

	// The API call will happen here
	// For example sake a simple delegation  is implemented
	p.subject.Book()
}

// Dummy Subject
type HotelBoutique struct{}

func (s *HotelBoutique) Book() {
	fmt.Println("Booking done on external site")
}

