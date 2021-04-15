package bookings

import (
	"time"

	"github.com/JoseFMP/resharmonics"
	"github.com/JoseFMP/resharmonics/utils"
)

type bookingsClient struct {
	resharmonics.Client
}

type BookingsClient interface {
	List(period utils.BookingPeriod, lastUpdated *time.Time, status []*BookingStatus) ([]*Booking, error)
	Get(bookingIdentified BookingIdentifier) (*Booking, error)
}

func Init(creds resharmonics.Credentials) (BookingsClient, error) {

	rhClient, errInitializingClient := resharmonics.Init(creds)
	if errInitializingClient != nil {
		return nil, errInitializingClient
	}
	return &bookingsClient{
		rhClient,
	}, nil
}
