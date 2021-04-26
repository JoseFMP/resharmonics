package invoices

type FinancialEntityID struct {
	Type         FinancialEntityIDType           `json:"type"`
	ID           int                             `json:"id"`
	ExternalID   string                          `json:"externalId"`
	ExportStatus FinancialIdentifierExportStatus `json:"exportStatus"`

	// ExportDate $date-time
	ExportDate string `json:"exportDate"`
}

type FinancialEntityIDType string

const FinancialEntityIDTypeFINANCE_ACCOUNT = FinancialEntityIDType("FINANCE_ACCOUNT")
const FinancialEntityIDTypeINVOICE = FinancialEntityIDType("INVOICE")
const FinancialEntityIDTypeCREDIT_NOTE = FinancialEntityIDType("CREDIT_NOTE")
const FinancialEntityIDTypePAYMENT = FinancialEntityIDType("PAYMENT")
const FinancialEntityIDTypeREFUND = FinancialEntityIDType("REFUND")
const FinancialEntityIDTypePURCHASE_INVOICE = FinancialEntityIDType("PURCHASE_INVOICE")
const FinancialEntityIDTypePURCHASE_CREDIT_NOTE = FinancialEntityIDType("PURCHASE_CREDIT_NOTE")

type FinancialIdentifierExportStatus string

const FinancialEntityIdExportStatusSKIPPED = FinancialIdentifierExportStatus("SKIPPED")
const FinancialEntityIdExportStatusCREATED = FinancialIdentifierExportStatus("CREATED")
const FinancialEntityIdExportStatusPENDING = FinancialIdentifierExportStatus("PENDING")
const FinancialEntityIdExportStatusSENT = FinancialIdentifierExportStatus("SENT")
const FinancialEntityIdExportStatusEXPORTED = FinancialIdentifierExportStatus("EXPORTED")
const FinancialEntityIdExportStatusFAILED = FinancialIdentifierExportStatus("FAILED")
const FinancialEntityIdExportStatusIMPORTED = FinancialIdentifierExportStatus("IMPORTED")
