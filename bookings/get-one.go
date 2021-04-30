package bookings

import (
	"encoding/json"
	"fmt"

	"github.com/JoseFMP/resharmonics/contact"
	"github.com/JoseFMP/resharmonics/invoices"
	"github.com/JoseFMP/resharmonics/utils"
)

func (clt *bookingsClient) Get(id Identifier) (*BookingS, error) {

	validationRes := validateGetOneParams(id)
	if validationRes != nil {
		return nil, validationRes
	}

	subPath := fmt.Sprintf("%s/%s", bookingsSubpath, string(id))
	payloadRes, errDoingReq := clt.DoGet(subPath, nil)

	if errDoingReq != nil {
		return nil, errDoingReq
	}

	rawBooking, errParsing := parseGetOneResponse(payloadRes)
	if errParsing != nil {
		return nil, errParsing
	}
	asBooking, errConverting := rawBooking.toBooking()

	if errConverting != nil {
		return nil, errConverting
	}

	return asBooking, nil
}

func validateGetOneParams(id Identifier) error {
	if string(id) == "" {
		return fmt.Errorf("Booking identifier is empty")
	}
	return nil
}

func parseGetOneResponse(payload []byte) (*RawBookingS, error) {
	var booking RawBookingS
	errUnmarshalling := json.Unmarshal(payload, &booking)
	if errUnmarshalling != nil {
		return nil, errUnmarshalling
	}

	return &booking, nil
}

func (bookingRaw *RawBookingS) toBooking() (*BookingS, error) {

	startDate, errParsingStartDate := utils.FromDateString(bookingRaw.Period.From)
	if errParsingStartDate != nil {
		return nil, errParsingStartDate
	}

	endDate, errParsingEndDate := utils.FromDateString(bookingRaw.Period.To)
	if errParsingEndDate != nil {
		return nil, errParsingEndDate
	}

	invoices := make([]*invoices.Invoice, len(bookingRaw.Invoices))
	for index, in := range bookingRaw.Invoices {
		invoices[index] = &in
	}

	result := BookingS{
		Reference:  bookingRaw.Reference,
		Identifier: bookingRaw.Id,
		Period: utils.BookingPeriod{
			From: startDate,
			To:   endDate,
		},
		Guest:    bookingRaw.Guest,
		Invoices: invoices,
	}
	return &result, nil
}

type RawBookingS struct {
	Reference    BookingReference   `json:"bookingReference"`
	Id           Identifier         `json:"bookingIdentifier"`
	MaxOccupancy int                `json:"maxOccupancy"`
	Guest        contact.Details    `json:"contactDetails"`
	Period       RawSinglePeriod    `json:"datePeriod"`
	Extras       []Extra            `json:"extras"`
	FloorSpace   string             `json:"floorSpace"`
	Invoices     []invoices.Invoice `json:"invoices"`
}

type RawSinglePeriod struct {
	From string `json:"dateFrom"`
	To   string `json:"dateTo"`
}

// BookingS is just a bit more parsed and less raw than BookingData. Otherwise just the sae
type BookingS struct {
	Reference  BookingReference    `json:"bookingReference"`
	Identifier Identifier          `json:"bookingIdentifier"`
	Period     utils.BookingPeriod `json:"period"`
	Guest      contact.Details     `json:"guest"`
	Invoices   []*invoices.Invoice `json:"invoices"`
}
