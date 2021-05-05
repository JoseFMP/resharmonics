package bookings

import (
	"encoding/json"
	"fmt"
)

type BookingStatus string

type allBookingStatuses struct {
	Confirmed  BookingStatus
	CheckedIn  BookingStatus
	CheckedOut BookingStatus
	Pending    BookingStatus
	Cancelled  BookingStatus
}

func (clt *bookingsClient) GetStatus(id Identifier) (BookingStatus, error) {

	targetURL := fmt.Sprintf("%s/%s/%s", bookingsSubpath, string(id), statusSubpath)

	respPayload, errDoingGet := clt.DoGet(targetURL, nil)
	if errDoingGet != nil {
		return BookingStatus(""), errDoingGet
	}

	var statusResponse StatusResponse
	errUnmarshalling := json.Unmarshal(respPayload, &statusResponse)
	if errUnmarshalling != nil {
		return BookingStatus(""), errUnmarshalling
	}

	return statusResponse.Status, nil
}

const statusSubpath = `status`

type StatusResponse struct {
	Status BookingStatus `json:"status"`
}
