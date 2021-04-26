package invoices

type FinanceAccount struct {
	Type          FinanceAccountType         `json:"type"`
	AccountNumber string                     `json:"accountNumber"`
	ID            *FinancialEntityIdentifier `json:"id"`
}

type FinancialEntityIdentifier struct {
}

type FinanceAccountType string

const FinancyAccountTypeCompany = FinanceAccountType("COMPANY")
const FinancyAccountTypeContact = FinanceAccountType("CONTACT")
