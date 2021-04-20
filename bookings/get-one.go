package bookings

import (
	"encoding/json"
	"fmt"

	"github.com/JoseFMP/resharmonics/contact"
	"github.com/JoseFMP/resharmonics/invoices"
	"github.com/JoseFMP/resharmonics/utils"
)

func (clt *bookingsClient) Get(id Identifier) (*Booking, error) {

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

func parseGetOneResponse(payload []byte) (*RawSingleBooking, error) {
	var booking RawSingleBooking
	errUnmarshalling := json.Unmarshal(payload, &booking)
	if errUnmarshalling != nil {
		return nil, errUnmarshalling
	}

	return &booking, nil
}

func (bookingRaw *RawSingleBooking) toBooking() (*Booking, error) {

	startDate, errParsingStartDate := utils.FromDateString(bookingRaw.Period.From)
	if errParsingStartDate != nil {
		return nil, errParsingStartDate
	}

	endDate, errParsingEndDate := utils.FromDateString(bookingRaw.Period.To)
	if errParsingEndDate != nil {
		return nil, errParsingEndDate
	}

	var invoicesPointer *[]*invoices.Invoice
	if bookingRaw.Invoices != nil {
		invoices := make([]*invoices.Invoice, len(bookingRaw.Invoices))
		for index, in := range bookingRaw.Invoices {
			invoices[index] = in.ToInvoice()
		}
		invoicesPointer = &invoices
	}
	result := Booking{
		Reference:  bookingRaw.Reference,
		Identifier: bookingRaw.Id,
		Period: utils.BookingPeriod{
			From: startDate,
			To:   endDate,
		},
		Guests:   []contact.Details{bookingRaw.Guest},
		Invoices: invoicesPointer,
	}
	return &result, nil
}

type RawSingleBooking struct {
	Reference    BookingReference      `json:"bookingReference"`
	Id           Identifier            `json:"bookingIdentifier"`
	MaxOccupancy int                   `json:"maxOccupancy"`
	Guest        contact.Details       `json:"contactDetails"`
	Period       RawSinglePeriod       `json:"datePeriod"`
	Extras       []Extra               `json:"extras"`
	FloorSpace   string                `json:"floorSpace"`
	Invoices     []invoices.RawInvoice `json:"invoices"`
}

type RawSinglePeriod struct {
	From string `json:"dateFrom"`
	To   string `json:"dateTo"`
}
