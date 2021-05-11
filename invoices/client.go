package invoices

import (
	"github.com/JoseFMP/resharmonics"
	"github.com/JoseFMP/resharmonics/utils"
)

type invoicesClient struct {
	resharmonics.Client
}

type InvoicesClient interface {
	List(period utils.BookingPeriod, org resharmonics.OrganizationID, pagination *utils.Pagination) ([]*Invoice, error)
	//Get(bookingIdentified Identifier) (*BookingS, error)
	//GetStatus(id Identifier) (BookingStatus, error)
}

func Init(creds resharmonics.Credentials, preAuthorize bool) (InvoicesClient, error) {

	rhClient, errInitializingClient := resharmonics.Init(creds, preAuthorize)
	if errInitializingClient != nil {
		return nil, errInitializingClient
	}
	return &invoicesClient{
		rhClient,
	}, nil
}
