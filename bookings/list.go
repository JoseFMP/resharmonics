package bookings

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/JoseFMP/resharmonics/contact"
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

func composeGetParams(period *utils.BookingPeriod, lastUpdated *time.Time, statusFilter []*BookingStatus) (map[string]string, error) {

	result := map[string]string{
		dateFromParamName: period.From.ToResharmonicsString(),
		dateToParamName:   period.To.ToResharmonicsString(),
	}

	if lastUpdated != nil {
		result[lastUpdatedParamName] = lastUpdated.Format(time.RFC3339)
	}

	if statusFilter != nil && len(statusFilter) > 0 {
		marshalled, errMarshallingStatuses := json.Marshal(statusFilter)

		if errMarshallingStatuses != nil {
			return nil, fmt.Errorf("Error marshalling statuses: %v", errMarshallingStatuses)
		}
		result[statusesParamName] = string(marshalled)
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
	Reference  BookingReference      `json:"bookingReference"`
	Identifier Identifier            `json:"bookingIdentifier"`
	Status     BookingStatus         `json:"status"`
	Period     utils.BookingPeriod   `json:"period"`
	Guests     []contact.Details     `json:"guests"`
	Property   property.PropertyData `json:"property"`
}
