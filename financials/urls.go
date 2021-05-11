package financials

import "fmt"

const financialsSubpath = `financials`

var invoicesSubpath = fmt.Sprintf(`%s/invoices`, financialsSubpath)
