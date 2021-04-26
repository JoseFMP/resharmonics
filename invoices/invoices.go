package invoices

import "github.com/JoseFMP/resharmonics/address"

type Invoice struct {
	// Number is unique
	Number         InvoiceNumber      `json:"invoiceNumber"`
	FinanceAccount *FinanceAccount    `json:"financeAccount"`
	ID             *FinancialEntityID `json:"id"`
	TotalNet       float64            `json:"totalNet"`
	TotalVat       float64            `json:"totalVat"`
	Currency       Currency           `json:"currencyCode"`

	// InvoiceDate $date, no time
	InvoiceDate string `json:"invoiceDate"`

	InvoiceDueDate string `json:"invoiceDueDate"`
}

type InvoiceNumber string

type InvoiceDetails struct {
	CustomerName string           `json:"customerName"`
	CompanyName  string           `json:"companyName"`
	Address      *address.Address `json:"address"`
}
