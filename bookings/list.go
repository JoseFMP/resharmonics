package bookings

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/JoseFMP/resharmonics/contact"
	"github.com/JoseFMP/resharmonics/invoices"
	"github.com/JoseFMP/resharmonics/property"
	"github.com/JoseFMP/resharmonics/utils"
)

func (clt *bookingsClient) List(period utils.BookingPeriod, lastUpdated *time.Time, statusesFilter []*BookingStatus) ([]*BookingL, error) {

	validationRes := validateListParams(&period, lastUpdated, statusesFilter)
	if validationRes != nil {
		return nil, validationRes
	}

	getParams, errPreparingGetParams := composeGetParams(&period, lastUpdated, statusesFilter)
	if errPreparingGetParams != nil {
		return nil, errPreparingGetParams
	}

	payloadRes, errDoingReq := clt.DoGet(bookingsSubpath, getParams)

	if errDoingReq != nil {
		return nil, errDoingReq
	}

	rawBookings, errParsing := parseListResponse(payloadRes)
	if errParsing != nil {
		return nil, errParsing
	}
	res := make([]*BookingL, 0)

	for _, rb := range rawBookings {
		if rb == nil {
			continue
		}
		asBooking, errConverting := rb.toBooking()
		if errConverting != nil {
			return nil, errConverting
		}
		res = append(res, asBooking)
	}

	return res, nil
}

func composeGetParams(period *utils.BookingPeriod, lastUpdated *time.Time, statusFilter []*BookingStatus) (map[string]interface{}, error) {

	result := map[string]interface{}{
		dateFromParamName: period.From.ToResharmonicsString(),
		dateToParamName:   period.To.ToResharmonicsString(),
	}

	if lastUpdated != nil {
		result[lastUpdatedParamName] = lastUpdated.Format(time.RFC3339)
	}

	if statusFilter != nil && len(statusFilter) > 0 {

		filterMapped := make([]string, len(statusFilter))

		for i, s := range statusFilter {
			filterMapped[i] = string(*s)
		}

		result[statusesParamName] = filterMapped
	}

	return result, nil
}

func validateListParams(period *utils.BookingPeriod, lastUpdated *time.Time, statuses []*BookingStatus) error {
	if period == nil {
		return fmt.Errorf("No period specified")
	}
	validationPeriod := period.Validate()
	if validationPeriod != nil {
		return validationPeriod
	}

	if lastUpdated != nil {

	}

	if statuses != nil && len(statuses) > 0 {
		for _, s := range statuses {
			if s == nil {
				return fmt.Errorf("Status is in filter is nil")
			}
		}
	}

	return nil
}

const bookingsSubpath = `bookings`

const dateFromParamName = `dateFrom`
const dateToParamName = `dateTo`
const lastUpdatedParamName = `latUpdated`
const statusesParamName = `statuses`

type BookingSearchResponse struct {
	Results []*BookingData `json:"results"`
}

func parseListResponse(payload []byte) ([]*BookingData, error) {
	var response BookingSearchResponse
	errUnmarshalling := json.Unmarshal(payload, &response)
	if errUnmarshalling != nil {
		return nil, errUnmarshalling
	}
	return response.Results, nil
}

// BookingL is just a bit more parsed and less raw than BookingData. Otherwise just the sae
type BookingL struct {
	Id        Identifier            `json:"bookingIdentifier"`
	Reference BookingReference      `json:"bookingReference"`
	Status    BookingStatus         `json:"status"`
	Period    utils.BookingPeriod   `json:"period"`
	Guests    []contact.Details     `json:"guests"`
	Property  property.PropertyData `json:"property"`
	Extras    []Extra               `json:"extras"`
}

// BookingData is the raw payload format as returned by the Resharmonics API endpoint /bookings
type BookingData struct {
	Id                 Identifier            `json:"bookingIdentifier"`
	Reference          BookingReference      `json:"bookingReference"`
	Status             BookingStatus         `json:"status"`
	StartDate          string                `json:"startDate"` // just date as 2005-01-01
	EndDate            string                `json:"endDate"`   // just date as 2005-01-01
	Guests             []contact.Details     `json:"guests"`
	Property           property.PropertyData `json:"property"`
	NightlyAverageRate float64               `json:"nightlyAverageRate"`
	Invoices           []invoices.Invoice    `json:"invoices"`
	Extras             []Extra               `json:"extras"`
	BookingAccountName string                `json:"bookingAccountName"`
	BillingAccountName string                `json:"billingAccountName"`
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
		Reference: bookingRaw.Reference,
		Id:        bookingRaw.Id,
		Status:    bookingRaw.Status,
		Period: utils.BookingPeriod{
			From: startDate,
			To:   endDate,
		},
		Guests:   bookingRaw.Guests,
		Property: bookingRaw.Property,
		Extras:   bookingRaw.Extras,
	}
	return &result, nil
}
