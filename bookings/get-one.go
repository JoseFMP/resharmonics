package bookings

import (
	"encoding/json"
	"fmt"

	"github.com/JoseFMP/resharmonics/contact"
	"github.com/JoseFMP/resharmonics/financials"
	"github.com/JoseFMP/resharmonics/property"
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

func (raw *RawBookingS) toBooking() (*BookingS, error) {

	startDate, errParsingStartDate := utils.FromDateString(raw.Period.From)
	if errParsingStartDate != nil {
		return nil, errParsingStartDate
	}

	endDate, errParsingEndDate := utils.FromDateString(raw.Period.To)
	if errParsingEndDate != nil {
		return nil, errParsingEndDate
	}

	invoices := make([]*financials.Invoice, len(raw.Invoices))
	for index, in := range raw.Invoices {
		invoices[index] = &in
	}

	result := BookingS{
		Reference:  raw.Reference,
		Identifier: raw.Id,
		Period: utils.BookingPeriod{
			From: startDate,
			To:   endDate,
		},
		ContactDetails: raw.Guest,
		Invoices:       invoices,

		Property: property.PropertyData{
			BuildingName: raw.BuildingName,
			MaxOccupancy: raw.MaxOccupancy,
			Address:      &raw.BuildingAddress,
		},
	}

	if raw.UnitName != nil {
		result.Property.Name = *raw.UnitName
	}
	if raw.FloorSpace != nil {
		result.Property.FloorSpace = *raw.FloorSpace
	}

	return &result, nil
}

type RawBookingS struct {
	Reference       BookingReference     `json:"bookingReference"`
	Id              Identifier           `json:"bookingIdentifier"`
	MaxOccupancy    int                  `json:"maxOccupancy"`
	Guest           contact.Details      `json:"contactDetails"`
	Period          RawSinglePeriod      `json:"datePeriod"`
	Extras          []Extra              `json:"extras"`
	FloorSpace      *string              `json:"floorSpace"`
	UnitName        *string              `json:"unitName"`
	UnitType        string               `json:"unitType"`
	Description     string               `json:"description"`
	Invoices        []financials.Invoice `json:"invoices"`
	BuildingAddress property.Address     `json:"buildingAddress"`
	BuildingName    string               `json:"buildingName"`
	Features        []property.Feature   `json:"features"`
	Currency        string               `json:"currencySymbol"`
}

type RawSinglePeriod struct {
	From string `json:"dateFrom"`
	To   string `json:"dateTo"`
}

// BookingS is just a bit more parsed and less raw than BookingData. Otherwise just the sae
type BookingS struct {
	Reference      BookingReference      `json:"bookingReference"`
	Identifier     Identifier            `json:"bookingIdentifier"`
	Period         utils.BookingPeriod   `json:"period"`
	ContactDetails contact.Details       `json:"contactDetails"`
	Invoices       []*financials.Invoice `json:"invoices"`
	Property       property.PropertyData `json:"property"`
}
