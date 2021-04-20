package bookings

import (
	"github.com/JoseFMP/resharmonics/utils"
)

type BookingStatus string

type allBookingStatuses struct {
	Confirmed  BookingStatus
	CheckedIn  BookingStatus
	CheckedOut BookingStatus
	Pending    BookingStatus
}

type Identifier string
type BookingReference string

func getAllBookingStatuses() *allBookingStatuses {

	return &allBookingStatuses{
		Confirmed:  "CONFIRMED",
		CheckedIn:  "CHECKED_IN",
		CheckedOut: "CHECKED_OUT",
		Pending:    "PENDING",
	}
}

func (bookingRaw *BookingData) toBooking() (*BookingL, error) {

	startDate, errParsingStartDate := utils.FromDateString(bookingRaw.StartDate)
	if errParsingStartDate != nil {
		return nil, errParsingStartDate
	}

	endDate, errParsingEndDate := utils.FromDateString(bookingRaw.EndDate)
	if errParsingEndDate != nil {
		return nil, errParsingEndDate
	}

	result := BookingL{
		Reference:  bookingRaw.Reference,
		Identifier: bookingRaw.Id,
		Status:     bookingRaw.Status,
		Period: utils.BookingPeriod{
			From: startDate,
			To:   endDate,
		},
		Guests:   bookingRaw.Guests,
		Property: bookingRaw.Property,
	}
	return &result, nil
}
