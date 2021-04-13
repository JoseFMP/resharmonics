package bookings

import (
	"github.com/JoseFMP/resharmonics/contact"
	"github.com/JoseFMP/resharmonics/property"
	"github.com/JoseFMP/resharmonics/utils"
)

// BookingData is the raw payload format as returned by the Resharmonics API
type BookingData struct {
	Identifier         BookingIdentifier     `json:"bookingIdentifier"`
	Reference          BookingReference      `json:"bookingReference"`
	Status             BookingStatus         `json:"status"`
	StartDate          string                `json:"startDate"` // just date as 2005-01-01
	EndDate            string                `json:"endDate"`   // just date as 2005-01-01
	Guests             []contact.Details     `json:"guests"`
	Property           property.PropertyData `json:"property"`
	NightlyAverageRate float64               `json:"NightlyAverageRate"`
}

// Booking is just a bit more parsed and less raw than BookingData. Otherwise just the sae
type Booking struct {
	Reference  BookingReference      `json:"bookingReference"`
	Identifier BookingIdentifier     `json:"bookingIdentifier"`
	Status     BookingStatus         `json:"status"`
	Period     utils.BookingPeriod   `json:"period"`
	Guests     []contact.Details     `json:"guests"`
	Property   property.PropertyData `json:"property"`
}

type BookingStatus string

type allBookingStatuses struct {
	Confirmed  BookingStatus
	CheckedIn  BookingStatus
	CheckedOut BookingStatus
	Pending    BookingStatus
}

type BookingIdentifier string
type BookingReference string

func getAllBookingStatuses() *allBookingStatuses {

	return &allBookingStatuses{
		Confirmed:  "CONFIRMED",
		CheckedIn:  "CHECKED_IN",
		CheckedOut: "CHECKED_OUT",
		Pending:    "PENDING",
	}
}

func (bookingRaw *BookingData) toBooking() (*Booking, error) {

	startDate, errParsingStartDate := utils.FromDateString(bookingRaw.StartDate)
	if errParsingStartDate != nil {
		return nil, errParsingStartDate
	}

	endDate, errParsingEndDate := utils.FromDateString(bookingRaw.EndDate)
	if errParsingEndDate != nil {
		return nil, errParsingEndDate
	}

	result := Booking{
		Reference:  bookingRaw.Reference,
		Identifier: bookingRaw.Identifier,
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
