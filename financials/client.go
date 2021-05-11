package financials

import (
	"github.com/JoseFMP/resharmonics"
	"github.com/JoseFMP/resharmonics/utils"
)

type financialsClient struct {
	resharmonics.Client
}

type FinancialsClient interface {
	ListInvoices(period utils.BookingPeriod, org resharmonics.OrganizationID, pagination *utils.Pagination) ([]*Invoice, error)
	GetInvoice(invoiceID FinancialEntityIDID) (*Invoice, error)
}

func Init(creds resharmonics.Credentials, preAuthorize bool) (FinancialsClient, error) {

	rhClient, errInitializingClient := resharmonics.Init(creds, preAuthorize)
	if errInitializingClient != nil {
		return nil, errInitializingClient
	}
	return &financialsClient{
		rhClient,
	}, nil
}
