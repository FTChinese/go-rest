package gorest

import "net/http"

// Pagination is used to calculate limit and offset parameter used int sql statement.
type Pagination struct {
	page  int64 // Which page is requesting data.
	Limit int64 // How many items per page.
}

// NewPagination creates a new Pagination instance.
// p is the page number, r is the rows to retrieve.
func NewPagination(p, limit int64) Pagination {
	if p < 1 {
		p = 1
	}

	return Pagination{
		page:  p,
		Limit: limit,
	}
}

// Offset calculate the offset for SQL.
func (p Pagination) Offset() int64 {
	return (p.page - 1) * p.Limit
}

// GetPagination extracts pagination information from query parameter
func GetPagination(req *http.Request) Pagination {
	page, _ := GetQueryParam(req, "page").ToInt()
	perPage, err := GetQueryParam(req, "per_page").ToInt()
	if err != nil {
		perPage = 20
	}

	return NewPagination(page, perPage)
}
