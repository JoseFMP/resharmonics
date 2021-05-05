package bookings

import "strings"

type Identifier string
type BookingReference string

func (original *BookingReference) AsCanonical() string {

	val := string(*original)

	splitted := strings.Split(val, "/")

	return splitted[0]
}

func GetAllBookingStatuses() *allBookingStatuses {

	return &allBookingStatuses{
		Confirmed:  "CONFIRMED",
		CheckedIn:  "CHECKED_IN",
		CheckedOut: "CHECKED_OUT",
		Pending:    "PENDING",
		Cancelled:  "CANCELLED",
	}
}
