package bookings

import "github.com/JoseFMP/resharmonics/utils"

type BookingRaw struct {
	BookingReference BookingIdentifier `json:"bookingReference"`
	Status           BookingStatus     `json:"status"`
	StartDate        string            `json:"startDate"`
	EndDate          string            `json:"endDate"`
}

type Booking struct {
	BookingReference BookingIdentifier  `json:"bookingReference"`
	Status           BookingStatus      `json:"status"`
	StartDate        *utils.BookingDate `json:"startDate"`
	EndDate          *utils.BookingDate `json:"endDate"`
}

type BookingStatus string

type allBookingStatuses struct {
	Confirmed  BookingStatus
	CheckedIn  BookingStatus
	CheckedOut BookingStatus
	Pending    BookingStatus
}

type BookingIdentifier string

func getAllBookingStatuses() *allBookingStatuses {

	return &allBookingStatuses{
		Confirmed:  "CONFIRMED",
		CheckedIn:  "CHECKED_IN",
		CheckedOut: "CHECKED_OUT",
		Pending:    "PENDING",
	}
}

func (bookingRaw *BookingRaw) toBooking() (*Booking, error) {

	startDate, errParsingStartDate := utils.FromDateString(bookingRaw.StartDate)
	if errParsingStartDate != nil {
		return nil, errParsingStartDate
	}

	endDate, errParsingEndDate := utils.FromDateString(bookingRaw.StartDate)
	if errParsingEndDate != nil {
		return nil, errParsingEndDate
	}

	result := Booking{
		BookingReference: bookingRaw.BookingReference,
		Status:           bookingRaw.Status,
		StartDate:        startDate,
		EndDate:          endDate,
	}
	return &result, nil
}
