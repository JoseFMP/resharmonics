package financials

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/JoseFMP/resharmonics"
	"github.com/JoseFMP/resharmonics/utils"
)

func (clt *financialsClient) ListInvoices(period utils.BookingPeriod, org resharmonics.OrganizationID, pagination *utils.Pagination) ([]*Invoice, error) {
	validationRes := validateListParams(period, org, pagination)
	if validationRes != nil {
		return nil, validationRes
	}

	getParams := composeGetParams(period, org, pagination)

	payloadRes, errDoingReq := clt.DoGet(invoicesSubpath, getParams)

	if errDoingReq != nil {
		return nil, errDoingReq
	}

	invoices, errParsing := parseListResponse(payloadRes)
	if errParsing != nil {
		return nil, errParsing
	}
	res := make([]*Invoice, 0)

	for _, inv := range invoices {
		if inv == nil {
			continue
		}
		errValidating := validateInvoice(inv)
		if errValidating != nil {
			log.Printf("Skipping invoice: %v", errValidating)
			continue
		}

		res = append(res, inv)
	}

	return res, nil
}

type ListInvoicesResponse struct {
	Invoices []*Invoice `json:"invoices"`
}

func parseListResponse(payload []byte) ([]*Invoice, error) {

	var parsedResp ListInvoicesResponse
	errUnmarshalling := json.Unmarshal(payload, &parsedResp)
	if errUnmarshalling != nil {
		return nil, errUnmarshalling
	}
	return parsedResp.Invoices, nil
}

const dateFromParamName = `dateFrom`
const dateToParamName = `dateTo`
const organsationIDParamName = `organisationId`

func composeGetParams(period utils.BookingPeriod, org resharmonics.OrganizationID, pagination *utils.Pagination) map[string]interface{} {

	result := make(map[string]interface{})

	dateFrom := period.From.ToResharmonicsString()
	dateTo := period.To.ToResharmonicsString()

	result[dateFromParamName] = dateFrom
	result[dateToParamName] = dateTo
	result[organsationIDParamName] = fmt.Sprintf("%d", int(org))

	if pagination != nil {
		result[utils.PageSizeGetParamName] = fmt.Sprintf("%d", pagination.PageSize)
		if pagination.Page != nil {
			result[utils.PageGetParamName] = fmt.Sprintf("%d", *pagination.Page)
		}
	}

	return result

}

func validateListParams(period utils.BookingPeriod, org resharmonics.OrganizationID, pagination *utils.Pagination) error {

	if int(org) == 0 {
		return fmt.Errorf("Organization not specified, zero provided")
	}

	if int(org) < 0 {
		return fmt.Errorf("Organization is negative, wrong")
	}

	if pagination != nil {
		if pagination.PageSize < 1 {
			return fmt.Errorf("Page size less than 1?")
		}
	}

	errValidatingPeriod := period.Validate()
	if errValidatingPeriod != nil {
		return errValidatingPeriod
	}

	return nil
}
