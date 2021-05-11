package financials

type FinanceAccount struct {
	Type          FinanceAccountType         `json:"type"`
	AccountNumber string                     `json:"accountNumber"`
	ID            *FinancialEntityIdentifier `json:"id"`
	Email         string                     `json:"emailAddress"`
	AccountName   string                     `json:"accountName"`
}

type FinancialEntityIdentifier struct {
}

type FinanceAccountType string

const FinancyAccountTypeCompany = FinanceAccountType("COMPANY")
const FinancyAccountTypeContact = FinanceAccountType("CONTACT")
