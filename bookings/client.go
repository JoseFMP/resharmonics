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
	List(period utils.BookingPeriod, lastUpdated *time.Time, status []*BookingStatus, pagination *utils.Pagination) ([]*BookingL, error)
	Get(bookingIdentified Identifier) (*BookingS, error)
	GetStatus(id Identifier) (BookingStatus, error)
}

func Init(creds resharmonics.Credentials, preAuthorize bool) (BookingsClient, error) {

	rhClient, errInitializingClient := resharmonics.Init(creds, preAuthorize)
	if errInitializingClient != nil {
		return nil, errInitializingClient
	}
	return &bookingsClient{
		rhClient,
	}, nil
}
