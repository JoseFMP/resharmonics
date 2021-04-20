package invoices

type Invoice struct {
	Number InvoiceNumber `json:"invoiceNumber"`
}

type RawInvoice struct {
	Number InvoiceNumber `json:"invoiceNumber"`
}

type InvoiceNumber string

func (raw *RawInvoice) ToInvoice() *Invoice {

	return &Invoice{
		Number: raw.Number,
	}
}
