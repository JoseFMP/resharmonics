package financials

import (
	"fmt"

	"github.com/JoseFMP/resharmonics/property"
)

type Invoice struct {
	FinanceAccount *FinanceAccount `json:"financeAccount"`
	// Number is unique
	Number   InvoiceNumber      `json:"invoiceNumber"`
	ID       *FinancialEntityID `json:"id"`
	TotalNet float64            `json:"totalNet"`
	TotalVat float64            `json:"totalTax"`
	Currency Currency           `json:"currencyCode"`

	// InvoiceDate $date, no time
	InvoiceDate string `json:"invoiceDate"`

	InvoiceDueDate string `json:"invoiceDueDate"`

	InvoiceReference string     `json:"invoiceReference,omitempty"`
	Items            []LineItem `json:"lineItems"`
}

type InvoiceNumber string

type InvoiceDetails struct {
	CustomerName string            `json:"customerName"`
	CompanyName  string            `json:"companyName"`
	Address      *property.Address `json:"address"`
}

func (inv *Invoice) Validate() error {
	if inv.ID == nil {
		return fmt.Errorf("No ID in this invoice")
	}
	if int((inv).ID.ID) < 0 {
		return fmt.Errorf("Invoice ID is invalid")
	}
	return nil
}

func (inv *Invoice) ExtractID() FinancialEntityIDID {
	return inv.ID.ID
}
