package financials

import (
	"encoding/json"
	"fmt"
)

func (clt *financialsClient) GetInvoice(invoiceID FinancialEntityIDID) (*Invoice, error) {
	validationRes := validateGetOneParams(invoiceID)
	if validationRes != nil {
		return nil, validationRes
	}

	targetSubpath := fmt.Sprintf("%s/%d", invoicesSubpath, int(invoiceID))

	payloadRes, errDoingReq := clt.DoGet(targetSubpath, nil)

	if errDoingReq != nil {
		return nil, errDoingReq
	}

	invoice, errParsing := parseGetOneResponse(payloadRes)
	if errParsing != nil {
		return nil, errParsing
	}

	if invoice != nil {
		errValidating := invoice.Validate()
		if errValidating != nil {
			return nil, errValidating
		}
	}
	return invoice, nil
}

func parseGetOneResponse(payload []byte) (*Invoice, error) {

	var parsedResp Invoice
	errUnmarshalling := json.Unmarshal(payload, &parsedResp)
	if errUnmarshalling != nil {
		return nil, errUnmarshalling
	}
	return &parsedResp, nil
}

func validateGetOneParams(invoiceID FinancialEntityIDID) error {

	if int(invoiceID) == 0 {
		return fmt.Errorf("Invoice ID looks empty, i.e. is zero")
	}

	if int(invoiceID) < 0 {
		return fmt.Errorf("Invoice ID looks not valid, i.e. is negative")
	}

	return nil
}
