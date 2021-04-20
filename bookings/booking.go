package bookings

import (
	"github.com/JoseFMP/resharmonics/contact"
	"github.com/JoseFMP/resharmonics/invoices"
	"github.com/JoseFMP/resharmonics/property"
	"github.com/JoseFMP/resharmonics/utils"
)

// Booking is just a bit more parsed and less raw than BookingData. Otherwise just the sae
type Booking struct {
	Reference  BookingReference       `json:"bookingReference"`
	Identifier Identifier             `json:"bookingIdentifier"`
	Status     *BookingStatus         `json:"status"`
	Period     utils.BookingPeriod    `json:"period"`
	Guests     []contact.Details      `json:"guests"`
	Property   *property.PropertyData `json:"property"`
	Invoices   *[]*invoices.Invoice   `json:"invoices"`
}

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
		Identifier: bookingRaw.Id,
		Status:     &bookingRaw.Status,
		Period: utils.BookingPeriod{
			From: startDate,
			To:   endDate,
		},
		Guests:   bookingRaw.Guests,
		Property: &bookingRaw.Property,
	}
	return &result, nil
}
