package main

import (
	"fmt"
	"strings"
	"time"
)

const reservationDateFormat = "2006-01-02"

type ReservationTime time.Time

func (t *ReservationTime) UnmarshalJSON(bytes []byte) error {
	rawT, err := time.Parse(reservationDateFormat, strings.Replace(
		string(bytes),
		"\"",
		"",
		-1,
	))

	if err != nil {
		return err
	}

	*t = ReservationTime(rawT)

	return nil
}

func (t *ReservationTime) MarshalJSON() ([]byte, error) {
	buf := fmt.Sprintf("\"%s\"", time.Time(*t).Format(reservationDateFormat))
	return []byte(buf), nil
}
