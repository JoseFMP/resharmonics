package bookings

import (
	"time"
)

func (clt *bookingsClient) List(from time.Time, to time.Time, lastUpdated *time.Time, status *ListBookingFilter) ([]*Booking, error) {

	return nil, nil
}
