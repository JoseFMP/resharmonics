package utils

type Pagination struct {
	PageSize int
	Page     *int
}

const PageSizeGetParamName = `pageSize`
const PageGetParamName = `page`
