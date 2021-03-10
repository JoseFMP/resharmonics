package bookings

type Booking struct {
	BookingReference BookingIdentifier `json:"bookingReference"`
	Status           BookingStatus     `json:"status"`
	StartDate        string            `json:"startDate"`
	EndDate          string            `json:"endDate"`
}

type BookingStatus string

type allBookingStatuses struct {
	Confirmed  BookingStatus
	CheckedIn  BookingStatus
	CheckedOut BookingStatus
	Pending    BookingStatus
}

type BookingIdentifier string
