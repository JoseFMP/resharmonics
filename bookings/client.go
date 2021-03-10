package bookings

import (
	"time"

	"github.com/JoseFMP/resharmonics"
)

type bookingsClient struct {
	resharmonics.Resharmonics
}

type Client interface {
	List(from time.Time, to time.Time, lastUpdated *time.Time, status *ListBookingFilter) ([]*Booking, error)
	Get(bookingIdentified BookingIdentifier) (*Booking, error)
}

type ListBookingFilter string

func Init(username string, password string) (Client, error) {

	rhClient, errInitializingClient := resharmonics.Init(username, password)
	if errInitializingClient != nil {
		return nil, errInitializingClient
	}
	return &bookingsClient{
		rhClient,
	}, nil
}
